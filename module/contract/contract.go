package contract

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

// BIFContractService ...
type BIFContractService interface {
	// CheckContractAddress 检测合约账户的有效性
	CheckContractAddress(bifContractCheckValidRequest request.BIFContractCheckValidRequest) response.BIFContractCheckValidResponse
	// ContractCreate 创建合约
	ContractCreate(bifContractCreateRequest request.BIFContractCreateRequest) response.BIFContractCreateResponse
	// GetContractInfo 查询合约代码
	GetContractInfo(bifContractGetInfoRequest request.BIFContractGetInfoRequest) response.BIFContractGetInfoResponse
	// GetContractAddress 根据交易Hash查询合约地址
	GetContractAddress(bifContractGetAddressRequest request.BIFContractGetAddressRequest) response.BIFContractGetAddressResponse
	// ContractQuery 合约查询接口
	ContractQuery(bifContractCallRequest request.BIFContractCallRequest) response.BIFContractCallResponse
	// ContractInvoke 合约调用
	ContractInvoke(bifContractInvokeRequest request.BIFContractInvokeRequest) response.BIFContractInvokeResponse
	// BatchContractInvoke 批量调用合约
	BatchContractInvoke(bifBatchContractInvokeRequest request.BIFBatchContractInvokeRequest) response.BIFContractInvokeResponse
}

// ContractService ...
type ContractService struct {
	url string
}

func GetContractInstance(url string) *ContractService {
	return &ContractService{
		url,
	}
}

func (cs *ContractService) GetContractInfo(r request.BIFContractGetInfoRequest) response.BIFContractGetInfoResponse {

	if cs.url == "" {
		return response.BIFContractGetInfoResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if !key.IsAddressValid(r.ContractAddress) {
		return response.BIFContractGetInfoResponse{
			BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
		}
	}

	return blockchain.GetContractInfo(cs.url, r.ContractAddress)
}

func (cs *ContractService) GetContractAddress(r request.BIFContractGetAddressRequest) response.BIFContractGetAddressResponse {

	if r.Hash == "" {
		return response.BIFContractGetAddressResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if len(r.Hash) != common.HASH_HEX_LENGTH {
		return response.BIFContractGetAddressResponse{
			BIFBaseResponse: exception.INVALID_HASH_ERROR,
		}
	}

	transactionService := blockchain.GetTransactionInstance(cs.url)
	var transactionGetInfoRequest request.BIFTransactionGetInfoRequest
	transactionGetInfoRequest.Hash = r.Hash
	transactionGetInfoResponse := transactionService.GetTransactionInfo(transactionGetInfoRequest)
	if transactionGetInfoResponse.ErrorCode != common.SUCCESS {
		return response.BIFContractGetAddressResponse{
			BIFBaseResponse: transactionGetInfoResponse.BIFBaseResponse,
		}
	}

	var contractAddressInfos []response.ContractAddressInfo
	for _, v := range transactionGetInfoResponse.Result.Transactions {
		var contractInfos []response.ContractAddressInfo
		err := json.Unmarshal([]byte(v.ErrorDesc), &contractInfos)
		if err != nil {
			return response.BIFContractGetAddressResponse{
				BIFBaseResponse: exception.SYSTEM_ERROR,
			}
		}
		contractAddressInfos = append(contractAddressInfos, contractInfos...)
	}

	return response.BIFContractGetAddressResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFContractGetAddressResult{
			ContractAddressInfos: contractAddressInfos,
		},
	}
}

func (cs *ContractService) CheckContractAddress(r request.BIFContractCheckValidRequest) response.BIFContractCheckValidResponse {

	if r.ContractAddress == "" {
		return response.BIFContractCheckValidResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}

	if !key.IsAddressValid(r.ContractAddress) {
		return response.BIFContractCheckValidResponse{
			BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
		}
	}

	var isValid bool
	contractGetInfoResponse := blockchain.GetContractInfo(cs.url, r.ContractAddress)
	if contractGetInfoResponse.BIFBaseResponse.ErrorCode != common.SUCCESS {
		isValid = false
	} else {
		isValid = true
	}

	return response.BIFContractCheckValidResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFContractCheckValidResult{
			IsValid: isValid,
		},
	}
}

func (cs *ContractService) ContractQuery(r request.BIFContractCallRequest) response.BIFContractCallResponse {

	if r.SourceAddress != "" && !key.IsAddressValid(r.SourceAddress) {
		return response.BIFContractCallResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.ContractAddress != "" && !key.IsAddressValid(r.ContractAddress) {
		return response.BIFContractCallResponse{
			BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
		}
	}
	if r.SourceAddress == r.ContractAddress {
		return response.BIFContractCallResponse{
			BIFBaseResponse: exception.SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	var resp response.BIFContractCallResponse
	if r.Type == common.CONTRACT_TYPE_EVM {
		contractCallResponse := callContractEvm(cs.url, r.SourceAddress, r.ContractAddress, common.CONTRACT_QUERY_OPT_TYPE, r.Input, r.GasPrice, r.FeeLimit)
		resp.BIFBaseResponse = contractCallResponse.BIFBaseResponse
		resp.Result = contractCallResponse.Result
		return resp
	}
	contractCallResponse := callContractJs(cs.url, r.SourceAddress, r.ContractAddress, common.CONTRACT_QUERY_OPT_TYPE, r.Input, r.GasPrice, r.FeeLimit)
	resp.BIFBaseResponse = contractCallResponse.BIFBaseResponse
	resp.Result = contractCallResponse.Result
	return resp
}

func callContractJs(url string, sourceAddress string, contractAddress string, optType int, input string, gasPrice int64, feeLimit int64) response.BIFContractCallJsResponse {
	if url == "" {
		return response.BIFContractCallJsResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}

	params := make(map[string]interface{})
	params["opt_type"] = optType
	params["fee_limit"] = feeLimit
	params["source_address"] = sourceAddress
	params["contract_address"] = contractAddress
	params["input"] = input
	params["gas_price"] = gasPrice
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return response.BIFContractCallJsResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	// call contract
	contractCallURL := common.ContractCallURL(url)
	dataByte, err := http.HttpPost(contractCallURL, paramsByte)
	if err != nil {
		return response.BIFContractCallJsResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFContractCallJsResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFContractCallJsResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func callContractEvm(url string, sourceAddress string, contractAddress string, optType int, input string, gasPrice int64, feeLimit int64) response.BIFContractCallEvmResponse {
	if url == "" {
		return response.BIFContractCallEvmResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}

	params := make(map[string]interface{})
	params["opt_type"] = optType
	params["fee_limit"] = feeLimit
	params["source_address"] = sourceAddress
	params["contract_address"] = contractAddress
	params["input"] = input
	params["gas_price"] = gasPrice
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return response.BIFContractCallEvmResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	// call contract
	contractCallURL := common.ContractCallURL(url)
	dataByte, err := http.HttpPost(contractCallURL, paramsByte)
	if err != nil {
		return response.BIFContractCallEvmResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFContractCallEvmResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFContractCallEvmResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (cs *ContractService) ContractInvoke(r request.BIFContractInvokeRequest) response.BIFContractInvokeResponse {

	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if !key.IsAddressValid(r.ContractAddress) {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
		}
	}
	//if r.BIFAmount == 0 || r.BIFAmount < common.INIT_ZERO {
	//	return response.BIFContractInvokeResponse{
	//		BIFBaseResponse: exception.INVALID_AMOUNT_ERROR,
	//	}
	//}
	if r.PrivateKey == "" {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 检查合约的有效性
	if !blockchain.CheckContractValid(cs.url, r.ContractAddress) {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR,
		}
	}

	// 广播交易
	transactionService := blockchain.GetTransactionInstance(cs.url)
	bifAccountActivateOperation := request.BIFContractInvokeOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.CONTRACT_INVOKE,
		},
		ContractAddress: r.ContractAddress,
		BifAmount:       r.BIFAmount,
		Input:           r.Input,
	}

	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        bifAccountActivateOperation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Metadata,
		SenderPrivateKey: r.PrivateKey,
	}

	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFContractInvokeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFContractInvokeResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

func (cs *ContractService) ContractCreate(r request.BIFContractCreateRequest) response.BIFContractCreateResponse {

	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if r.InitBalance != 0 && r.InitBalance <= common.INIT_ZERO {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}
	if r.Payload == "" {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.PAYLOAD_EMPTY_ERROR,
		}
	}
	if r.InitBalance != 0 && r.InitBalance <= common.INIT_ZERO {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 广播交易
	transactionService := blockchain.GetTransactionInstance(cs.url)
	contractCreateOperation := request.BIFContractCreateOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.CONTRACT_CREATE,
		},
		InitBalance: r.InitBalance,
		Type:        r.Type,
		Payload:     r.Payload,
		InitInput:   r.InitInput,
	}

	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        contractCreateOperation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Metadata,
		SenderPrivateKey: r.PrivateKey,
	}

	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFContractCreateResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFContractCreateResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFContractInvokeResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

func (cs *ContractService) BatchContractInvoke(r request.BIFBatchContractInvokeRequest) response.BIFContractInvokeResponse {
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}

	var contractAddressArray []string
	for _, opt := range r.Operations {
		contractAddressArray = append(contractAddressArray, opt.ContractAddress)
		if opt.BifAmount < common.INIT_ZERO {
			return response.BIFContractInvokeResponse{
				BIFBaseResponse: exception.INVALID_AMOUNT_ERROR,
			}
		}
	}

	for _, v := range contractAddressArray {
		if !key.IsAddressValid(v) {
			return response.BIFContractInvokeResponse{
				BIFBaseResponse: exception.INVALID_CONTRACTADDRESS_ERROR,
			}
		}
		if !blockchain.CheckContractValid(cs.url, v) {
			return response.BIFContractInvokeResponse{
				BIFBaseResponse: exception.CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR,
			}
		}

	}

	if r.PrivateKey == "" {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}
	if r.GasPrice < common.INIT_ZERO {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: exception.INVALID_GASPRICE_ERROR,
		}
	}

	// 广播交易
	bifRadioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        r.Operations,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}

	transactionService := blockchain.GetTransactionInstance(cs.url)
	radioTransactionResponse := transactionService.RadioTransaction(bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFContractInvokeResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFContractInvokeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFContractInvokeResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}
