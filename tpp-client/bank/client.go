package bank

import "crypto/tls"

type OpenBankingClient interface {
	GetAccessToken(requestId string) (*AccessToken, error)
	CreatePaymentAuthRequest(payer *PaymentAuthRequest, access *AccessToken, beneficiary *BeneficiaryDetails) (*PaymentAuthResponse, error)
	SubmitPayment(data *PaymentAuthGranted) (*SubmitPaymentResponse, error)
	GetPaymentStatus(data *PaymentAuthGranted) (*PaymentStatusResponse, error)
}

// PaymentAuthRequest the details for the payer
type PaymentAuthRequest struct {
	RequestId     string
	InstitutionId string
	Amount        string
	Payer         BeneficiaryDetails
}

// BeneficiaryDetails are the details of a Payer or Beneficiary
type BeneficiaryDetails struct {
	SortCode      string
	AccountNumber string
	Name          string
}

// AccessToken wraps on Oauth2 token
type AccessToken struct {
	Token     string
	ExpiresIn int
}

// OauthClientCreds represents the various data items on Oauth client might use to identify itself
type OauthClientCreds struct {
	ClientId       string
	ClientSecret   string
	TransportCert  tls.Certificate
	SigningCert    tls.Certificate
	RedirectionUrl string
}

type PaymentAuthResponse struct {
	RequestId string
	Url       string
	ConsentId string
}

type PaymentAuthGranted struct {
	ConsentId   string
	RequestId   string
	ConsentCode string
}

type SubmitPaymentResponse struct {
	RequestId    string
	ConsentCode  string
	ConsentToken string
	PaymentId    string
}

type PaymentStatusResponse struct {
	RequestId string
	Status    string
}
