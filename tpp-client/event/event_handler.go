package event

//import (
//	"encoding/json"
//	"github.com/sgerogia/hello-stablecoin/tpp-client/bank"
//	"github.com/sgerogia/hello-stablecoin/tpp-client/contract"
//	"github.com/sgerogia/hello-stablecoin/tpp-client/encrypt"
//)
//
//type EventServiceImpl struct {
//	eventFilter *contract.ProvableGBPFilterer
//	keyPair     *encrypt.KeyPair
//	bankClient  *bank.OpenBankingClient
//}
//
//type MintRequestPayload struct {
//	institutionId *string `json:"institutionId"`
//	sortCode      *string `json:"sortCode"`
//	accountNumber *string `json:"accountNumber"`
//	name          *string `json:"name"`
//}
//
//func NewEventService(
//	_eventFilter *contract.ProvableGBPFilterer,
//	_keyPair *encrypt.KeyPair,
//	_bankClient *bank.OpenBankingClient) *EventServiceImpl {
//	return &EventServiceImpl{eventFilter: _eventFilter, keyPair: _keyPair, bankClient: _bankClient}
//}
//
//func (handler *EventServiceImpl) ExtractMintRequest(request *contract.ProvableGBPMintRequest) (*MintRequest, error) {
//
//	// encrypted data
//	var box encrypt.EthSigUtilBox
//	if err := json.Unmarshal(request.EncryptedData, &box); err != nil {
//		return nil, err
//	}
//	// decrypt
//	decr, err := handler.keyPair.Decrypt(&box)
//	if err != nil {
//		return nil, err
//	}
//	// recreate payload
//	var payload MintRequestPayload
//	if err = json.Unmarshal(decr, &payload); err != nil {
//		return nil, err
//	}
//	return &MintRequest{envelope: request, payload: &payload}, nil
//}
//
//func (handler *EventServiceImpl) ProcessMintRequest(request *MintRequest) (*bank.PaymentAuthResponse, error) {
//	requestId := string(request.envelope.RequestId[:])
//	token, err := (*handler.bankClient).GetAccessToken(requestId)
//	if err != nil {
//		return nil, err
//	}
//	var pAuthReq = request.toPaymentAuthRequest(token)
//	resp, err := (*handler.bankClient).CreatePaymentAuthRequest(&pAuthReq)
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (handler *EventServiceImpl) ExtractAuthGranted(request *contract.ProvableGBPAuthGranted) (*AuthGranted, error) {
//
//	// encrypted data
//	var box encrypt.EthSigUtilBox
//	if err := json.Unmarshal(request.GrantEncryptedData, &box); err != nil {
//		return nil, err
//	}
//	// decrypt
//	decr, err := handler.keyPair.Decrypt(&box)
//	if err != nil {
//		return nil, err
//	}
//	// recreate payload
//	var payload AuthGrantedPayload
//	if err = json.Unmarshal(decr, &payload); err != nil {
//		return nil, err
//	}
//	return &AuthGranted{envelope: request, payload: &payload}, nil
//}
//
//func (handler *EventServiceImpl) ProcessAuthGranted(request *AuthGranted) error {
//	return nil
//}
