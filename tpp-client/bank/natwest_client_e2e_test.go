package bank_test

import (
	"github.com/google/uuid"
	"github.com/sgerogia/hello-stablecoin/tpp-client/bank"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"
)

const INSTITUTION = "natwest-sandbox"
const AMOUNT = "1"

// TestNatwestClientE2E tests the end-to-end flow of the Natwest client.
// It requires the following env vars to be set, otherwise execution is skipped:
// - NATWEST_SANDBOX_CLIENT_ID: the client id of the Natwest sandbox app
// - NATWEST_SANDBOX_CLIENT_SECRET: the client secret of the Natwest sandbox app
// - NATWEST_SANDBOX_REDIRECT_URL: the redirect url of the Natwest sandbox app
// - NATWEST_SANDBOX_CUSTOMER_USERNAME: the username of the customer to use for the consent flow (format: CUSTOMER_ID_OR_CUSTOMER_NUMBER@domainof.your.app)
func TestNatwestSandboxClient_E2E(t *testing.T) {
	// skip if env vars are not there
	if !isEnvVarSet("NATWEST_SANDBOX_CLIENT_ID") ||
		!isEnvVarSet("NATWEST_SANDBOX_CLIENT_SECRET") ||
		!isEnvVarSet("NATWEST_SANDBOX_REDIRECT_URL") ||
		!isEnvVarSet("NATWEST_SANDBOX_CUSTOMER_USERNAME") {
		t.Skip("Skipping test, env vars not set")
	}

	// arrange
	creds := bank.OauthClientCreds{
		ClientId:       os.Getenv("NATWEST_SANDBOX_CLIENT_ID"),
		ClientSecret:   os.Getenv("NATWEST_SANDBOX_CLIENT_SECRET"),
		RedirectionUrl: os.Getenv("NATWEST_SANDBOX_REDIRECT_URL"),
	}
	l := zap.NewExample().Sugar()
	client := bank.NewNatwestSandboxClient(30, &creds, l)
	reqId := uuid.New().String()
	payer := bank.BeneficiaryDetails{
		Name:          "John Doe",
		SortCode:      "500000",
		AccountNumber: "12345601",
	}
	receiver := bank.BeneficiaryDetails{
		Name:          "ProvableGBP Limited",
		SortCode:      "500000",
		AccountNumber: "87654301",
	}
	customerUsername := os.Getenv("NATWEST_SANDBOX_CUSTOMER_USERNAME")

	// --- act & assert ---

	// 1) get access token
	accessToken, err := client.GetAccessToken(reqId)
	require.NoError(t, err)
	assert.NotEmpty(t, accessToken.Token)
	assert.Positive(t, accessToken.ExpiresIn)

	// 2) create payment auth request
	paymentAuthRequest := bank.PaymentAuthRequest{
		RequestId:     reqId,
		InstitutionId: INSTITUTION,
		Amount:        AMOUNT,
		Payer:         payer,
	}
	authResp, err := client.CreatePaymentAuthRequest(&paymentAuthRequest, accessToken, &receiver)
	require.NoError(t, err)
	assert.NotEmpty(t, authResp.Url)
	assert.NotEmpty(t, authResp.RequestId)

	// 3) approve consent
	authGranted, err := client.ApproveConsent(authResp, customerUsername)
	require.NoError(t, err)
	assert.NotEmpty(t, authGranted.ConsentCode)
	assert.NotEmpty(t, authGranted.RequestId)

	// 4) submit payment
	paymResp, err := client.SubmitPayment(authGranted, &paymentAuthRequest, &receiver)
	require.NoError(t, err)
	assert.NotEmpty(t, paymResp.RequestId)
	assert.NotEmpty(t, paymResp.PaymentId)
	assert.NotEmpty(t, paymResp.ConsentToken)
	assert.NotEmpty(t, paymResp.ConsentCode)

	// 5) check payment
	// repeat 10 times with 1 sec sleep until status is "AcceptedSettlementCompleted"
	n := 0
	for n < 10 {
		checkResp, err := client.GetPaymentStatus(paymResp)
		require.NoError(t, err)
		if checkResp.Status == "AcceptedSettlementCompleted" {
			break
		}
		time.Sleep(1 * time.Second)
		n++
	}
	assert.NotEqual(t, 10, n, "Payment not settled after 10 seconds")
}

func isEnvVarSet(e string) bool {
	return os.Getenv(e) != ""
}
