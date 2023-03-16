package account

import (
	"encoding/json"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/exception"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/blockchain"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"
)

// BIFAccountService account interface
type BIFAccountService interface {
	// CreateAccount 创建账户
	CreateAccount(bifCreateAccountRequest request.BIFCreateAccountRequest) response.BIFCreateAccountResponse
	// GetAccount 获取指定的账户信息
	GetAccount(bifAccountGetInfoRequest request.BIFAccountGetInfoRequest) response.BIFAccountGetInfoResponse
	// GetNonce 获取指定账户的nonce值
	GetNonce(bifAccountGetNonceRequest request.BIFAccountGetNonceRequest) response.BIFAccountGetNonceResponse
	// GetAccountBalance 获取指定账户的星火令的余额
	GetAccountBalance(bifAccountGetBalanceRequest request.BIFAccountGetBalanceRequest) response.BIFAccountGetBalanceResponse
	// SetMetadatas 设置metadata
	SetMetadatas(bifAccountSetMetadataRequest request.BIFAccountSetMetadatasRequest) response.BIFAccountSetMetadatasResponse
	// GetAccountMetadatas 获取指定账户的metadata信息
	GetAccountMetadatas(bifAccountGetMetadataRequest request.BIFAccountGetMetadatasRequest) response.BIFAccountGetMetadatasResponse
	// SetPrivilege 设置权限
	SetPrivilege(bifAccountSetPrivilegeRequest request.BIFAccountSetPrivilegeRequest) response.BIFAccountSetPrivilegeResponse
	// GetAccountPriv 获取账户权限
	GetAccountPriv(bifAccountPrivRequest request.BIFAccountPrivRequest) response.BIFAccountPrivResponse
}

// AccountService ...
type AccountService struct {
	url string
}

func GetAccountInstance(url string) *AccountService {
	return &AccountService{
		url,
	}
}

func (as *AccountService) CreateAccount(r request.BIFCreateAccountRequest) response.BIFCreateAccountResponse {

	if r.SenderAddress == "" || r.DestAddress == "" {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if !key.IsAddressValid(r.DestAddress) {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_DESTADDRESS_ERROR,
		}
	}
	if r.InitBalance < common.INIT_ZERO {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}

	if r.PrivateKey == "" {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}

	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 广播交易
	transactionService := blockchain.GetTransactionInstance(as.url)
	bifAccountActivateOperation := request.BIFAccountActivateOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.ACCOUNT_ACTIVATE,
		},
		DestAddress: r.DestAddress,
		InitBalance: r.InitBalance,
	}

	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        bifAccountActivateOperation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}

	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFCreateAccountResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFAccountCreateAccountResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

func (as *AccountService) GetAccount(r request.BIFAccountGetInfoRequest) response.BIFAccountGetInfoResponse {

	if r.Address == "" {
		return response.BIFAccountGetInfoResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.Address) {
		return response.BIFAccountGetInfoResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}

	accountGetInfoUrl := common.AccountGetInfoURL(as.url, r.Address)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFAccountGetInfoResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFAccountGetInfoResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFAccountGetInfoResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (as *AccountService) GetNonce(r request.BIFAccountGetNonceRequest) response.BIFAccountGetNonceResponse {

	if r.Address == "" {
		return response.BIFAccountGetNonceResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.Address) {
		return response.BIFAccountGetNonceResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}

	accountGetInfoUrl := common.AccountGetInfoURL(as.url, r.Address)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFAccountGetNonceResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}
	var res response.BIFAccountGetNonceResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFAccountGetNonceResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (as *AccountService) GetAccountBalance(r request.BIFAccountGetBalanceRequest) response.BIFAccountGetBalanceResponse {

	if r.Address == "" {
		return response.BIFAccountGetBalanceResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.Address) {
		return response.BIFAccountGetBalanceResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if as.url == "" {
		return response.BIFAccountGetBalanceResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	accountGetInfoUrl := common.AccountGetInfoURL(as.url, r.Address)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFAccountGetBalanceResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}
	var res response.BIFAccountGetBalanceResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFAccountGetBalanceResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (as *AccountService) SetMetadatas(r request.BIFAccountSetMetadatasRequest) response.BIFAccountSetMetadatasResponse {
	if r.SenderAddress == "" || r.PrivateKey == "" {
		return response.BIFAccountSetMetadatasResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFAccountSetMetadatasResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if r.Key == "" || len(r.Key) > common.METADATA_KEY_MAX {
		return response.BIFAccountSetMetadatasResponse{
			BIFBaseResponse: exception.INVALID_DATAKEY_ERROR,
		}
	}
	if r.Value == "" || len(r.Value) > common.METADATA_VALUE_MAX {
		return response.BIFAccountSetMetadatasResponse{
			BIFBaseResponse: exception.INVALID_DATAVALUE_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 广播交易
	transactionService := blockchain.GetTransactionInstance(as.url)
	bifAccountSetMetadataOperation := request.BIFAccountSetMetadataOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.ACCOUNT_SET_METADATA,
		},
		Key:        r.Key,
		Value:      r.Value,
		Version:    r.Version,
		DeleteFlag: r.DeleteFlag,
	}
	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        bifAccountSetMetadataOperation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}
	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFAccountSetMetadatasResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFAccountSetMetadatasResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFAccountSetMetadataResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

func (as *AccountService) GetAccountMetadatas(r request.BIFAccountGetMetadatasRequest) response.BIFAccountGetMetadatasResponse {
	if r.Address == "" {
		return response.BIFAccountGetMetadatasResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.Address) {
		return response.BIFAccountGetMetadatasResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if !(r.Key == "") && (len(r.Key) > common.METADATA_KEY_MAX || len(r.Key) < common.METADATA_KEY_MIN) {
		return response.BIFAccountGetMetadatasResponse{
			BIFBaseResponse: exception.INVALID_DATAKEY_ERROR,
		}
	}

	accountGetMetadataUrl := common.AccountGetMetadataURL(as.url, r.Address, r.Key)
	dataByte, err := http.HttpGet(accountGetMetadataUrl)
	if err != nil {
		return response.BIFAccountGetMetadatasResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFAccountGetMetadatasResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFAccountGetMetadatasResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (as *AccountService) GetAccountPriv(r request.BIFAccountPrivRequest) response.BIFAccountPrivResponse {
	if r.Address == "" {
		return response.BIFAccountPrivResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.Address) {
		return response.BIFAccountPrivResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}

	accountGetInfoUrl := common.AccountGetInfoURL(as.url, r.Address)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFAccountPrivResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFAccountPrivResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFAccountPrivResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (as *AccountService) SetPrivilege(r request.BIFAccountSetPrivilegeRequest) response.BIFAccountSetPrivilegeResponse {
	if r.SenderAddress == "" {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}

	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.GasPrice < common.INIT_ZERO {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: exception.INVALID_GAS_AMOUNT_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 广播交易
	transactionService := blockchain.GetTransactionInstance(as.url)
	bifAccountSetPrivilegeOperation := request.BIFAccountSetPrivilegeOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.ACCOUNT_SET_PRIVILEGE,
		},
		Signers:        r.Signers,
		TxThreshold:    r.TxThreshold,
		MasterWeight:   r.MasterWeight,
		TypeThresholds: r.TypeThresholds,
	}
	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        bifAccountSetPrivilegeOperation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}
	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)

	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFAccountSetPrivilegeResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFAccountSetPrivilegeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFAccountSetPrivilegeResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}
