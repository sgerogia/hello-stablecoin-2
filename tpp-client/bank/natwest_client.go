package bank

import (
	"encoding/json"
	"errors"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
	tmpl "text/template"
	"time"
)

const OAUTH_TOKEN = "https://ob.sandbox.natwest.com/token"
const CREATE_PAYMENT_CONSENT = "https://ob.sandbox.natwest.com/open-banking/v3.1/pisp/domestic-payment-consents"
const CREATE_AUTH_URL = "https://api.sandbox.natwest.com/authorize"
const EXECUTE_PAYMENT = "https://api.sandbox.natwest.com/open-banking/v3.1/pisp/domestic-payments"
const CHECK_PAYMENT = "https://api.sandbox.natwest.com/open-banking/v3.1/pisp/domestic-payments/"

type NatwestSandboxClient struct {
	client           *resty.Client
	noRedirectClient *resty.Client
	clientCreds      *OauthClientCreds
	l                *zap.SugaredLogger
	consentTmpl      *tmpl.Template
	paymentTmpl      *tmpl.Template
}

func NewNatwestSandboxClient(
	timeout int,
	creds *OauthClientCreds,
	_l *zap.SugaredLogger) *NatwestSandboxClient {

	clRedir := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	clNoRedir := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &NatwestSandboxClient{
		clientCreds:      creds,
		client:           resty.NewWithClient(&clRedir).SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)),
		noRedirectClient: resty.NewWithClient(&clNoRedir).SetRedirectPolicy(resty.NoRedirectPolicy()),
		l:                _l,
		consentTmpl:      tmpl.Must(tmpl.New("consent").Parse(CONSENT_PAYLOAD)),
		paymentTmpl:      tmpl.Must(tmpl.New("payment").Parse(PAYMENT_PAYLOAD)),
	}
}

// GetAccessToken generates an Oauth2 token based on the client's credentials
func (c *NatwestSandboxClient) GetAccessToken(requestId string) (*AccessToken, error) {

	c.l.Infow("Access token request",
		"reqId", requestId)
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=payments",
			c.clientCreds.ClientId,
			c.clientCreds.ClientSecret)).
		Post(OAUTH_TOKEN)

	if err != nil {
		return nil, err
	}

	c.l.Debugw("Access token response",
		"reqId", requestId,
		"status", resp.StatusCode())

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed access token request. Status %d", resp.StatusCode()))
	}

	var atResp AccessTokenResponse
	if err = json.Unmarshal(resp.Body(), &atResp); err != nil {
		return nil, err
	}

	c.l.Debugw("Access token",
		"reqId", requestId,
		"accessToken", atResp.AccessToken,
		"expires", atResp.ExpiresIn)

	var at = AccessToken{
		Token:     atResp.AccessToken,
		ExpiresIn: atResp.ExpiresIn,
	}
	return &at, nil
}

// CreatePaymentAuthRequest generates a consent authorisation URL for the given payer and beneficiary details.
func (c *NatwestSandboxClient) CreatePaymentAuthRequest(
	authRequest *PaymentAuthRequest,
	access *AccessToken,
	beneficiary *BeneficiaryDetails) (*PaymentAuthResponse, error) {

	c.l.Infow("Payment auth request",
		"reqId", authRequest.RequestId)

	// build the payload
	var p strings.Builder
	d := consentData{
		AuthRequest: authRequest,
		Beneficiary: beneficiary,
	}
	if err := c.consentTmpl.Execute(&p, d); err != nil {
		return nil, err
	}

	// make the call
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+access.Token).
		SetHeader("x-jws-signature", "IGNORED_DUE_TO_REDUCED_SECURITY").
		SetHeader("x-idempotency-key", uuid.New().String()).
		SetBody(p.String()).
		Post(CREATE_PAYMENT_CONSENT)

	if err != nil {
		return nil, err
	}

	c.l.Debugw("Payment auth response",
		"reqId", authRequest.RequestId,
		"status", resp.StatusCode())

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("Failed payment auth request. Status %d", resp.StatusCode()))
	}

	// extract consent
	var pauthResp map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &pauthResp); err != nil {
		return nil, err
	}

	data := pauthResp["Data"].(map[string]interface{})
	consent := data["ConsentId"].(string)
	c.l.Debugw("Consent received",
		"reqId", authRequest.RequestId,
		"consent", consent)

	// generate authorisation URL
	resp, err = c.noRedirectClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+access.Token).
		SetHeader("x-jws-signature", "IGNORED_DUE_TO_REDUCED_SECURITY").
		SetHeader("x-idempotency-key", uuid.New().String()).
		SetQueryParam("client_id", c.clientCreds.ClientId).
		SetQueryParam("response_type", "code id_token").
		SetQueryParam("scope", "openid payments").
		SetQueryParam("redirect_uri", "https://display-parameters.com/").
		SetQueryParam("request", consent).
		Get(CREATE_AUTH_URL)

	if err != nil && !strings.Contains(err.Error(), "auto redirect is disabled") {
		return nil, err
	}

	c.l.Debugw("Create auth URL response",
		"reqId", authRequest.RequestId,
		"status", resp.StatusCode())

	if resp.StatusCode() != http.StatusFound {
		return nil, errors.New(fmt.Sprintf("Failed create auth URL request. Status %d", resp.StatusCode()))
	}

	location := resp.Header().Get("location")

	return &PaymentAuthResponse{RequestId: authRequest.RequestId, Url: location, ConsentId: consent}, nil
}

// SubmitPayment exchanges the approved consent code for a token and submits the authorised payment to the bank
func (c *NatwestSandboxClient) SubmitPayment(
	authGranted *PaymentAuthGranted,
	paymentAuthRequest *PaymentAuthRequest,
	beneficiary *BeneficiaryDetails) (*SubmitPaymentResponse, error) {

	c.l.Debugw("Exchange consent code",
		"requestId", authGranted.RequestId)

	// 1) exchange consent code for token
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(fmt.Sprintf("client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code&code=%s",
			c.clientCreds.ClientId,
			c.clientCreds.ClientSecret,
			c.clientCreds.RedirectionUrl,
			authGranted.ConsentCode)).
		Post(OAUTH_TOKEN)

	if err != nil {
		return nil, err
	}

	c.l.Debugw("Exchange consent code response",
		"reqId", paymentAuthRequest.RequestId,
		"status", resp.StatusCode())

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed exchange consent code. Status %d. Body %s", resp.StatusCode(), resp.Body()))
	}

	// extract acc. token
	var exchResp map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &exchResp); err != nil {
		return nil, err
	}
	accessToken := exchResp["access_token"].(string)

	// build the payload
	var p strings.Builder
	d := paymentData{
		Consent:            authGranted.ConsentId,
		PaymentAuthRequest: paymentAuthRequest,
		Beneficiary:        beneficiary,
	}
	if err := c.paymentTmpl.Execute(&p, d); err != nil {
		return nil, err
	}

	// 2) submit the payment
	resp, err = c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("x-jws-signature", "IGNORED_DUE_TO_REDUCED_SECURITY").
		SetHeader("x-idempotency-key", uuid.New().String()).
		SetBody(p.String()).
		Post(EXECUTE_PAYMENT)

	if err != nil {
		return nil, err
	}

	c.l.Debugw("Submit payment response",
		"reqId", paymentAuthRequest.RequestId,
		"status", resp.StatusCode())

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("Failed submit payment. Status %d. Body %s", resp.StatusCode(), resp.Body()))
	}

	// extract payment id
	var pResp map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &pResp); err != nil {
		return nil, err
	}
	data := pResp["Data"].(map[string]interface{})
	paymentId := data["DomesticPaymentId"].(string)
	c.l.Debugw("Payment id received",
		"reqId", paymentAuthRequest.RequestId,
		"paymentId", paymentId)

	return &SubmitPaymentResponse{
		RequestId:    paymentAuthRequest.RequestId,
		ConsentCode:  authGranted.ConsentCode,
		ConsentToken: accessToken,
		PaymentId:    paymentId,
	}, nil
}

// GetPaymentStatus retrieves the status of a submitted payment
func (c *NatwestSandboxClient) GetPaymentStatus(data *SubmitPaymentResponse) (*PaymentStatusResponse, error) {

	c.l.Debugw("Check payment status",
		"requestId", data.RequestId)

	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+data.ConsentToken).
		Get(CHECK_PAYMENT + data.PaymentId)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed check payment. Status %d", resp.StatusCode()))
	}

	var pResp map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &pResp); err != nil {
		return nil, err
	}
	d := pResp["Data"].(map[string]interface{})
	paymentStatus := d["Status"].(string)

	return &PaymentStatusResponse{
		RequestId: data.RequestId,
		Status:    paymentStatus,
	}, nil
}

// ApproveConsent utility function to approve the consent request. Only used for testing.
func (c *NatwestSandboxClient) ApproveConsent(data *PaymentAuthResponse, username string) (*PaymentAuthGranted, error) {

	c.l.Debugw("Approve consent",
		"requestId", data.RequestId)

	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParam("client_id", c.clientCreds.ClientId).
		SetQueryParam("client_secret", c.clientCreds.ClientSecret).
		SetQueryParam("response_type", "code id_token").
		SetQueryParam("redirect_uri", c.clientCreds.RedirectionUrl).
		SetQueryParam("scope", "openid payments").
		SetQueryParam("request", data.ConsentId).
		SetQueryParam("authorization_mode", "AUTO_POSTMAN").
		SetQueryParam("authorization_result", "APPROVED").
		SetQueryParam("authorization_username", username).
		Get(CREATE_AUTH_URL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Failed approve consent. Status %d. Body %s", resp.StatusCode(), resp.Body()))
	}

	var pResp map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &pResp); err != nil {
		return nil, err
	}
	uri := pResp["redirectUri"].(string)
	// parse the consent code from the redirect uri
	re := regexp.MustCompile(`.*?\#code\=([\w|-]+)\&.+`)
	rs := re.FindStringSubmatch(uri)
	if len(rs) != 2 {
		return nil, errors.New("Failed to parse consent code. Redirect uri: " + uri)
	}
	consentCode := rs[1]

	return &PaymentAuthGranted{
		ConsentId:   data.ConsentId,
		RequestId:   data.RequestId,
		ConsentCode: consentCode,
	}, nil
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type consentData struct {
	AuthRequest *PaymentAuthRequest
	Beneficiary *BeneficiaryDetails
}

const CONSENT_PAYLOAD = `
{
	"Data": {
		"Initiation": {
			"InstructionIdentification": "instr-identification",
			"EndToEndIdentification": "e2e-identification",
			"DebtorAccount": {
				"SchemeName": "SortCodeAccountNumber",
				"Identification": "{{.AuthRequest.Payer.SortCode}}{{.AuthRequest.Payer.AccountNumber}}",
				"Name": "{{.AuthRequest.Payer.Name}}"
			},
			"InstructedAmount": {
				"Amount": "{{.AuthRequest.Amount}}",
				"Currency": "GBP"
			},
			"CreditorAccount": {
				"SchemeName": "SortCodeAccountNumber",
				"Identification": "{{.Beneficiary.SortCode}}{{.Beneficiary.AccountNumber}}",
				"Name": "{{.Beneficiary.Name}}"
			},
			"RemittanceInformation": {
				"Unstructured": "Provable GBP mint",
				"Reference": "Provable GBP mint"
			}
		}
	},
	"Risk": {
		"PaymentContextCode": "Services",
		"MerchantCategoryCode": null,
		"MerchantCustomerIdentification": null,
		"DeliveryAddress": null
	}
}
`

type paymentData struct {
	Consent            string
	PaymentAuthRequest *PaymentAuthRequest
	Beneficiary        *BeneficiaryDetails
}

const PAYMENT_PAYLOAD = `
{
	"Data": {
		"ConsentId": "{{.Consent}}",
		"Initiation": {
			"InstructionIdentification": "instr-identification",
			"EndToEndIdentification": "e2e-identification",
			"DebtorAccount": {
				"SchemeName": "SortCodeAccountNumber",
				"Identification": "{{.PaymentAuthRequest.Payer.SortCode}}{{.PaymentAuthRequest.Payer.AccountNumber}}",
				"Name": "{{.PaymentAuthRequest.Payer.Name}}"
			},
			"InstructedAmount": {
				"Amount": "{{.PaymentAuthRequest.Amount}}",
				"Currency": "GBP"
			},
			"CreditorAccount": {
				"SchemeName": "SortCodeAccountNumber",
				"Identification": "{{.Beneficiary.SortCode}}{{.Beneficiary.AccountNumber}}",
				"Name": "{{.Beneficiary.Name}}"
			},
			"RemittanceInformation": {
				"Unstructured": "Provable GBP mint",
				"Reference": "Provable GBP mint"
			}
		}
	},
	"Risk": {
		"PaymentContextCode": "Services",
		"MerchantCategoryCode": null,
		"MerchantCustomerIdentification": null,
		"DeliveryAddress": null
	}
}
`
