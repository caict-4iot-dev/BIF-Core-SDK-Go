package blockchain

import (
	"encoding/json"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/exception"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/proto"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"
	"math"
	"strconv"
)

// OperationService ...
type OperationService struct {
	url string
}

func GetOperationInstance(url string) *OperationService {
	return &OperationService{
		url,
	}
}

// GetOperations ...
func (os *OperationService) GetOperations(operation interface{}, transSourceAddress string) ([]*proto.Operation, response.BIFBaseResponse) {

	var operations []*proto.Operation

	switch operation.(type) {
	case request.BIFAccountActivateOperation:
		operationData, ok := operation.(request.BIFAccountActivateOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		if operationData.SourceAddress == transSourceAddress && transSourceAddress != "" {
			return operations, exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
		}
		operationResData := os.ActivateOperation(operationData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFAccountSetMetadataOperation:
		operationData, ok := operation.(request.BIFAccountSetMetadataOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.SetMetadataOperation(operationData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFAccountSetPrivilegeOperation:
		operationsData, ok := operation.(request.BIFAccountSetPrivilegeOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.SetPrivilegeOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFGasSendOperation:
		operationsData, ok := operation.(request.BIFGasSendOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.GasSendOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFContractCreateOperation:
		operationsData, ok := operation.(request.BIFContractCreateOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.ContractCreateOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFContractInvokeOperation:
		operationsData, ok := operation.(request.BIFContractInvokeOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.ContractInvokeOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case []request.BIFContractInvokeOperation:
		operationsData, ok := operation.([]request.BIFContractInvokeOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		for _, v := range operationsData {
			operationResData := os.ContractInvokeOperation(v)
			if operationResData.ErrorCode != 0 {
				return operations, operationResData.BIFBaseResponse
			}
			operations = append(operations, operationResData.Result.Operation)
		}
	case request.BIFPrivateContractCreateOperation:
		operationsData, ok := operation.(request.BIFPrivateContractCreateOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.PrivateContractCreateOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	case request.BIFPrivateContractCallOperation:
		operationsData, ok := operation.(request.BIFPrivateContractCallOperation)
		if !ok {
			return operations, exception.OPERATIONS_ONE_ERROR
		}
		operationResData := os.PrivateContractCallOperation(operationsData)
		if operationResData.ErrorCode != 0 {
			return operations, operationResData.BIFBaseResponse
		}
		operations = append(operations, operationResData.Result.Operation)
	default:
		return operations, exception.OPERATIONS_ONE_ERROR
	}

	return operations, exception.SUCCESS
}

func (os *OperationService) ActivateOperation(r request.BIFAccountActivateOperation) response.AccountActivateResponse {
	if !(r.SourceAddress == "") && !key.IsAddressValid(r.SourceAddress) {
		return response.AccountActivateResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}

	if !key.IsAddressValid(r.DestAddress) {
		return response.AccountActivateResponse{
			BIFBaseResponse: exception.INVALID_DESTADDRESS_ERROR,
		}
	}

	if r.SourceAddress == r.DestAddress && r.SourceAddress != "" {
		return response.AccountActivateResponse{
			BIFBaseResponse: exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR,
		}
	}
	if r.InitBalance <= 0 {
		return response.AccountActivateResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}

	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_CREATE_ACCOUNT,
		CreateAccount: &proto.OperationCreateAccount{
			DestAddress: r.DestAddress,
			Priv: &proto.AccountPrivilege{
				MasterWeight: common.INIT_ONE,
				Thresholds: &proto.AccountThreshold{
					TxThreshold: common.INIT_ONE,
				},
			},
			InitBalance: r.InitBalance,
		},
	}

	return response.AccountActivateResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) GasSendOperation(r request.BIFGasSendOperation) response.BIFGasSendOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFGasSendOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if !key.IsAddressValid(r.DestAddress) {
		return response.BIFGasSendOperationResponse{
			BIFBaseResponse: exception.INVALID_DESTADDRESS_ERROR,
		}
	}
	if r.SourceAddress != "" && r.SourceAddress == r.DestAddress {
		return response.BIFGasSendOperationResponse{
			BIFBaseResponse: exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR,
		}
	}
	if r.Amount <= common.INIT_ZERO {
		return response.BIFGasSendOperationResponse{
			BIFBaseResponse: exception.INVALID_GAS_AMOUNT_ERROR,
		}
	}

	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_PAY_COIN,
		PayCoin: &proto.OperationPayCoin{
			DestAddress: r.DestAddress,
			Amount:      r.Amount,
		},
	}

	return response.BIFGasSendOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) SetPrivilegeOperation(r request.BIFAccountSetPrivilegeOperation) response.BIFAccountSetPrivilegeOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFAccountSetPrivilegeOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.SourceAddress != "" {
		masterWeightInt, err := strconv.ParseInt(r.MasterWeight, 10, 64)
		if err != nil || masterWeightInt < common.INIT_ZERO_L || masterWeightInt > math.MaxUint32 {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_MASTERWEIGHT_ERROR,
			}
		}
	}

	for i := range r.Signers {
		if !key.IsAddressValid(r.Signers[i].Address) {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_SIGNER_ADDRESS_ERROR,
			}
		}
		if r.Signers[i].Weight > math.MaxUint32 || r.Signers[i].Weight < common.INIT_ZERO_L {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_SIGNER_WEIGHT_ERROR,
			}
		}
	}
	if r.TxThreshold != "" {
		txThresholdInt, err := strconv.ParseInt(r.TxThreshold, 10, 64)
		if err != nil || txThresholdInt < common.INIT_ZERO_L {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_TX_THRESHOLD_ERROR,
			}
		}
	}
	for i := range r.TypeThresholds {
		if r.TypeThresholds[i].Type > 100 || r.TypeThresholds[i].Type <= 0 {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_TYPETHRESHOLD_TYPE_ERROR,
			}
		}
		if r.TypeThresholds[i].Threshold < 0 {
			return response.BIFAccountSetPrivilegeOperationResponse{
				BIFBaseResponse: exception.INVALID_TYPE_THRESHOLD_ERROR,
			}
		}
	}
	signers := make([]*proto.Signer, len(r.Signers))
	for _, v := range r.Signers {
		var signer *proto.Signer
		signer.Address = v.Address
		signer.Weight = v.Weight
		signers = append(signers, signer)
	}
	TypeThresholds := make([]*proto.OperationTypeThreshold, len(r.TypeThresholds))
	for _, v := range r.TypeThresholds {
		var typeThreshold proto.OperationTypeThreshold
		typeThreshold.Threshold = v.Threshold
		typeThreshold.Type = (proto.Operation_Type)(v.Type)
	}
	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_SET_PRIVILEGE,
		SetPrivilege: &proto.OperationSetPrivilege{
			MasterWeight:   r.MasterWeight,
			Signers:        signers,
			TxThreshold:    r.TxThreshold,
			TypeThresholds: TypeThresholds,
		},
	}

	return response.BIFAccountSetPrivilegeOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFAccountSetPrivilegeOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) SetMetadataOperation(r request.BIFAccountSetMetadataOperation) response.BIFAccountSetMetadataOperationResponse {

	if !(r.SourceAddress == "") && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFAccountSetMetadataOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}

	if r.Key == "" || len(r.Key) > common.METADATA_KEY_MAX {
		return response.BIFAccountSetMetadataOperationResponse{
			BIFBaseResponse: exception.INVALID_DATAKEY_ERROR,
		}
	}

	if r.Value == "" || len(r.Value) > common.METADATA_VALUE_MAX {
		return response.BIFAccountSetMetadataOperationResponse{
			BIFBaseResponse: exception.INVALID_DATAVALUE_ERROR,
		}
	}

	if r.Version < common.VERSION {
		return response.BIFAccountSetMetadataOperationResponse{
			BIFBaseResponse: exception.INVALID_DATAVERSION_ERROR,
		}
	}

	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_SET_METADATA,
		SetMetadata: &proto.OperationSetMetadata{
			Key:        r.Key,
			Value:      r.Value,
			Version:    r.Version,
			DeleteFlag: r.DeleteFlag,
		},
	}

	return response.BIFAccountSetMetadataOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) ContractCreateOperation(r request.BIFContractCreateOperation) response.BIFContractCreateOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFContractCreateOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.InitBalance <= common.INIT_ZERO {
		return response.BIFContractCreateOperationResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}
	if r.Type < 0 {
		return response.BIFContractCreateOperationResponse{
			BIFBaseResponse: exception.INVALID_CONTRACT_TYPE_ERROR,
		}
	}
	if r.Payload == "" {
		return response.BIFContractCreateOperationResponse{
			BIFBaseResponse: exception.PAYLOAD_EMPTY_ERROR,
		}
	}

	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_CREATE_ACCOUNT,
		CreateAccount: &proto.OperationCreateAccount{
			InitBalance: r.InitBalance,
			InitInput:   r.InitInput,
			Contract: &proto.Contract{
				Payload: r.Payload,
				Type:    proto.Contract_ContractType(r.Type),
			},
			Priv: &proto.AccountPrivilege{
				MasterWeight: common.INIT_ZERO,
				Thresholds: &proto.AccountThreshold{
					TxThreshold: common.INIT_ONE,
				},
			},
		},
	}

	return response.BIFContractCreateOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) ContractInvokeOperation(r request.BIFContractInvokeOperation) response.BIFContractInvokeOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFContractInvokeOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.ContractAddress != "" && !key.IsAddressValid(r.ContractAddress) {
		return response.BIFContractInvokeOperationResponse{
			BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
		}
	}
	if r.SourceAddress == r.ContractAddress {
		return response.BIFContractInvokeOperationResponse{
			BIFBaseResponse: exception.SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR,
		}
	}
	//if r.BifAmount <= common.INIT_ZERO {
	//	return response.BIFContractInvokeOperationResponse{
	//		BIFBaseResponse: exception.INVALID_AMOUNT_ERROR,
	//	}
	//}
	if !checkContractValid(os.url, r.ContractAddress) {
		return response.BIFContractInvokeOperationResponse{
			BIFBaseResponse: exception.CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR,
		}
	}

	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_PAY_COIN,
		PayCoin: &proto.OperationPayCoin{
			Input:       r.Input,
			Amount:      r.BifAmount,
			DestAddress: r.ContractAddress,
		},
	}

	return response.BIFContractInvokeOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func checkContractValid(url string, contractAddress string) bool {
	res := GetContractInfo(url, contractAddress)
	if res.ErrorCode != common.SUCCESS {
		return false
	}
	return true
}

func (os *OperationService) PrivateContractCreateOperation(r request.BIFPrivateContractCreateOperation) response.BIFPrivateContractCreateOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFPrivateContractCreateOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.Type < 0 {
		return response.BIFPrivateContractCreateOperationResponse{
			BIFBaseResponse: exception.INVALID_CONTRACT_TYPE_ERROR,
		}
	}
	if r.Payload == "" {
		return response.BIFPrivateContractCreateOperationResponse{
			BIFBaseResponse: exception.PAYLOAD_EMPTY_ERROR,
		}
	}
	privateTransactionSendRequest := request.BIFPrivateTransactionSendRequest{
		From:    r.From,
		Payload: r.Payload,
		To:      r.To,
	}
	privateTransactionService := GetPrivateTransactionInstance(os.url)
	privateTransactionSendResponse := privateTransactionService.Send(privateTransactionSendRequest)
	if privateTransactionSendResponse.ErrorCode != common.SUCCESS {
		return response.BIFPrivateContractCreateOperationResponse{
			BIFBaseResponse: privateTransactionSendResponse.BIFBaseResponse,
		}
	}

	payload := privateTransactionSendResponse.PriTxHash
	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_CREATE_PRIVATE_CONTRACT,
		CreatePrivateContract: &proto.OperationCreatePrivateContract{
			InitInput: r.InitInput,
			Contract: &proto.Contract{
				Payload: payload,
				Type:    proto.Contract_ContractType(r.Type),
			},
		},
	}
	return response.BIFPrivateContractCreateOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func (os *OperationService) PrivateContractCallOperation(r request.BIFPrivateContractCallOperation) response.PrivateContractCallOperationResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.PrivateContractCallOperationResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.Type < 0 {
		return response.PrivateContractCallOperationResponse{
			BIFBaseResponse: exception.INVALID_CONTRACT_TYPE_ERROR,
		}
	}
	privateTransactionSendRequest := request.BIFPrivateTransactionSendRequest{
		From:        r.From,
		Payload:     r.Input,
		To:          r.To,
		DestAddress: r.DestAddress,
	}
	privateTransactionService := GetPrivateTransactionInstance(os.url)
	privateTransactionSendResponse := privateTransactionService.Send(privateTransactionSendRequest)
	if privateTransactionSendResponse.ErrorCode != common.SUCCESS {
		return response.PrivateContractCallOperationResponse{
			BIFBaseResponse: privateTransactionSendResponse.BIFBaseResponse,
		}
	}
	input := privateTransactionSendResponse.PriTxHash
	operation := &proto.Operation{
		SourceAddress: r.SourceAddress,
		Metadata:      []byte(r.Metadata),
		Type:          proto.Operation_CALL_PRIVATE_CONTRACT,
		CallPrivateContract: &proto.OperationCallPrivateContract{
			DestAddress: r.DestAddress,
			Input:       input,
		},
	}

	return response.PrivateContractCallOperationResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFBaseOperationResult{
			Operation: operation,
		},
	}
}

func GetContractInfo(url string, address string) response.BIFContractGetInfoResponse {
	contractGetInfoURL := common.AccountGetInfoURL(url, address)
	dataByte, err := http.HttpGet(contractGetInfoURL)
	if err != nil {
		return response.BIFContractGetInfoResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFContractGetInfoResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFContractGetInfoResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	if res.Result.Contract.Payload == "" {
		return response.BIFContractGetInfoResponse{
			BIFBaseResponse: exception.CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR,
		}
	}

	return res
}
