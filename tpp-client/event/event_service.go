package event

//import (
//	"github.com/sgerogia/hello-stablecoin/tpp-client/bank"
//	"github.com/sgerogia/hello-stablecoin/tpp-client/contract"
//)
//
//type EventService interface {
//	ExtractMintRequest(request *contract.ProvableGBPMintRequest) (*MintRequest, error)
//	ProcessMintRequest(request *MintRequest) error
//	ExtractAuthGranted(request *contract.ProvableGBPAuthGranted) (*AuthGranted, error)
//	ProcessAuthGranted(request *AuthGranted) error
//}
//
//type MintRequest struct {
//	envelope *contract.ProvableGBPMintRequest
//	payload  *MintRequestPayload
//}
//
//type AuthGrantedPayload struct {
//	token *string
//}
//
//type AuthGranted struct {
//	envelope *contract.ProvableGBPAuthGranted
//	payload  *AuthGrantedPayload
//}
//
//func (mintReq *MintRequest) toPaymentAuthRequest(token *string) bank.PaymentAuthRequest {
//
//	tmpReq := string(mintReq.envelope.RequestId[:])
//	return bank.PaymentAuthRequest{
//		Token:         token,
//		RequestId:     &tmpReq,
//		InstitutionId: mintReq.payload.institutionId,
//		SortCode:      mintReq.payload.sortCode,
//		AccountNumber: mintReq.payload.accountNumber,
//		Name:          mintReq.payload.name,
//	}
//}
