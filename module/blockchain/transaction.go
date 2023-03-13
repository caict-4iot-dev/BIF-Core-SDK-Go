package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/exception"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/proto"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/hash"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"
	protobuf "github.com/golang/protobuf/proto"
)

// BIFTransactionService ...
type BIFTransactionService interface {
	// GasSend 发送交易
	GasSend(r request.BIFTransactionGasSendRequest) response.BIFTransactionGasSendResponse
	// PrivateContractCreate 私有化交易-合约创建
	PrivateContractCreate(r request.BIFTransactionPrivateContractCreateRequest) response.BIFTransactionPrivateContractCreateResponse
	// PrivateContractCall 私有化交易-合约调用
	PrivateContractCall(r request.BIFTransactionPrivateContractCallRequest) response.BIFTransactionPrivateContractCallResponse
	// GetTransactionInfo 根据交易hash查询交易
	GetTransactionInfo(r request.BIFTransactionGetInfoRequest) response.BIFTransactionGetInfoResponse
	// EvaluateFee 交易的费用评估
	EvaluateFee(r request.BIFTransactionEvaluateFeeRequest) response.BIFTransactionEvaluateFeeResponse
	// BIFSubmit 交易提交
	BIFSubmit(r request.BIFTransactionSubmitRequest) response.BIFTransactionSubmitResponse
	// GetTxCacheSize 获取交易池中交易条数
	GetTxCacheSize() response.BIFTransactionGetTxCacheSizeResponse
	// GetTxCacheData 获取交易池交易数据
	GetTxCacheData(r request.BIFTransactionCacheRequest) response.BIFTransactionCacheResponse
	// BatchGasSend 批量转账
	BatchGasSend(r request.BIFBatchGasSendRequest) response.BIFTransactionGasSendResponse
}

// TransactionService ...
type TransactionService struct {
	url string
}

func GetTransactionInstance(url string) *TransactionService {
	return &TransactionService{
		url,
	}
}

func (ts *TransactionService) GasSend(r request.BIFTransactionGasSendRequest) response.BIFTransactionGasSendResponse {
	if ts.url == "" {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.SenderAddress == "" || r.DestAddress == "" {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if !key.IsAddressValid(r.DestAddress) {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_DESTADDRESS_ERROR,
		}
	}
	if r.Amount < common.INIT_ZERO {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_GAS_AMOUNT_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	operation := request.BIFGasSendOperation{
		DestAddress: r.DestAddress,
		Amount:      r.Amount,
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.GAS_SEND,
		},
	}

	radioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        operation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}
	// 广播交易
	radioTransactionResponse := ts.RadioTransaction(radioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFTransactionGasSendResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionGasSendResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

// RadioTransaction 广播交易
func (ts *TransactionService) RadioTransaction(r request.BIFRadioTransactionRequest) response.BIFRadioTransactionResponse {

	// 一、获取交易发起的账号nonce值
	accountGetInfoUrl := common.AccountGetInfoURL(ts.url, r.SenderAddress)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}
	var nonceResponse response.BIFAccountGetNonceResponse
	err = json.Unmarshal(dataByte, &nonceResponse)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	if nonceResponse.ErrorCode != common.SUCCESS {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: nonceResponse.BIFBaseResponse,
		}
	}
	nonce := nonceResponse.Result.Nonce

	// 二、构建操作、序列化交易
	// 初始化请求参数 BIFTransactionSerializeRequest
	serializeRequest := request.BIFTransactionSerializeRequest{
		SourceAddress: r.SenderAddress,
		Nonce:         nonce + 1,
		GasPrice:      r.GasPrice,
		FeeLimit:      r.FeeLimit,
		Operation:     r.Operation,
		CeilLedgerSeq: r.CeilLedgerSeq,
		Metadata:      r.Remarks,
	}
	// BIFTransactionSerializeResponse
	serializeResponse := ts.BIFSerializable(serializeRequest)
	if serializeResponse.ErrorCode != common.SUCCESS {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: serializeResponse.BIFBaseResponse,
		}
	}
	transactionBlob := serializeResponse.Result.TransactionBlob
	blob, err := hex.DecodeString(transactionBlob)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	// 三、签名
	signData, err := key.Sign([]byte(r.SenderPrivateKey), blob)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	publicKey, err := key.GetEncPublicKey([]byte(r.SenderPrivateKey))
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	// 四、提交交易 BIFTransactionSubmitRequest
	submitRequest := request.BIFTransactionSubmitRequest{
		Serialization: transactionBlob,
		SignData:      hex.EncodeToString(signData),
		PublicKey:     publicKey,
	}

	// 调用bifSubmit接口 BIFTransactionSubmitResponse
	transactionSubmitResponse := ts.BIFSubmit(submitRequest)
	if transactionSubmitResponse.ErrorCode != common.SUCCESS {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: transactionSubmitResponse.BIFBaseResponse,
		}
	}

	return response.BIFRadioTransactionResponse{
		BIFBaseResponse: transactionSubmitResponse.BIFBaseResponse,
		Result: response.BIFRadioTransactionResult{
			Hash: transactionSubmitResponse.Result.Hash,
		},
	}
}

// BIFSubmit 交易提交
func (ts *TransactionService) BIFSubmit(r request.BIFTransactionSubmitRequest) response.BIFTransactionSubmitResponse {
	if ts.url == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.SignData == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SIGNDATA_NULL_ERROR,
		}
	}
	if r.PublicKey == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.PUBLICKEY_NULL_ERROR,
		}
	}
	if r.Serialization == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
		}
	}

	var transactionSubmit request.TransactionSubmit
	transactionSubmit.TransactionBlob = r.Serialization
	var signature request.Signature
	signature.SignData = r.SignData
	signature.PublicKey = r.PublicKey
	transactionSubmit.Signatures = append(transactionSubmit.Signatures, signature)

	var transactionSubmitRequest request.TransactionSubmitRequest
	transactionSubmitRequest.Items = append(transactionSubmitRequest.Items, transactionSubmit)
	transactionRequest, err := json.Marshal(transactionSubmitRequest)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	submitURL := common.TransactionSubmitURL(ts.url)
	dataByte, err := http.HttpPost(submitURL, transactionRequest)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.TransactionSubmitResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	if res.Results[0].ErrorCode != common.SUCCESS {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: response.BIFBaseResponse{
				ErrorCode: res.Results[0].ErrorCode,
				ErrorDesc: res.Results[0].ErrorDesc,
			},
		}
	}

	return response.BIFTransactionSubmitResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionSubmitResult{
			Hash: res.Results[0].Hash,
		},
	}
}

// BIFSerializable 交易序列化
func (ts *TransactionService) BIFSerializable(r request.BIFTransactionSerializeRequest) response.BIFTransactionSerializeResponse {

	if !key.IsAddressValid(r.SourceAddress) {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.Nonce <= 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_NONCE_ERROR,
		}
	}
	if r.CeilLedgerSeq < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_CEILLEDGERSEQ_ERROR,
		}
	}
	if r.GasPrice < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_GASPRICE_ERROR,
		}
	}
	if r.FeeLimit < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}

	if r.Operation == nil {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.OPERATIONS_EMPTY_ERROR,
		}
	}

	operationService := GetOperationInstance(ts.url)
	operations, bifBaseResponse := operationService.GetOperations(r.Operation, r.SourceAddress)
	if bifBaseResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: bifBaseResponse,
		}
	}
	if r.CeilLedgerSeq < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_CEILLEDGERSEQ_ERROR,
		}
	}
	var seq int64 = 0
	if r.CeilLedgerSeq > 0 {
		blockService := GetBlockInstance(ts.url)
		blockGetNumberResponse := blockService.GetBlockNumber()
		if blockGetNumberResponse.ErrorCode != common.SUCCESS {
			return response.BIFTransactionSerializeResponse{
				BIFBaseResponse: blockGetNumberResponse.BIFBaseResponse,
			}
		}

		seq = r.CeilLedgerSeq + blockGetNumberResponse.Result.Header.BlockNumber
	}
	transaction := proto.Transaction{
		SourceAddress: r.SourceAddress,
		Nonce:         r.Nonce,
		CeilLedgerSeq: seq,
		FeeLimit:      r.FeeLimit,
		GasPrice:      r.GasPrice,
		Metadata:      []byte(r.Metadata),
		Operations:    operations,
	}
	blobByte, err := protobuf.Marshal(&transaction)
	if err != nil {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	blob := hex.EncodeToString(blobByte)

	return response.BIFTransactionSerializeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionSerializeResult{
			TransactionBlob: blob,
			Hash:            string(hash.GenerateHashHex(blobByte, hash.SHA256)),
		},
	}
}

func (ts *TransactionService) GetTransactionInfo(r request.BIFTransactionGetInfoRequest) response.BIFTransactionGetInfoResponse {

	if ts.url == "" {
		return response.BIFTransactionGetInfoResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.Hash == "" {
		return response.BIFTransactionGetInfoResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if len(r.Hash) != common.HASH_HEX_LENGTH {
		return response.BIFTransactionGetInfoResponse{
			BIFBaseResponse: exception.INVALID_HASH_ERROR,
		}
	}
	getInfoUrl := common.TransactionGetInfoURL(ts.url, r.Hash)
	dataByte, err := http.HttpGet(getInfoUrl)
	if err != nil {
		return response.BIFTransactionGetInfoResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFTransactionGetInfoResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFTransactionGetInfoResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

// Deprecated
func (ts *TransactionService) PrivateContractCreate(r request.BIFTransactionPrivateContractCreateRequest) response.BIFTransactionPrivateContractCreateResponse {

	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFTransactionPrivateContractCreateResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFTransactionPrivateContractCreateResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.Type < 0 {
		return response.BIFTransactionPrivateContractCreateResponse{
			BIFBaseResponse: exception.INVALID_CONTRACT_TYPE_ERROR,
		}
	}
	if r.Payload == "" {
		return response.BIFTransactionPrivateContractCreateResponse{
			BIFBaseResponse: exception.PAYLOAD_EMPTY_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	operation := request.BIFPrivateContractCreateOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.PRIVATE_CONTRACT_CREATE,
		},
		Type:      r.Type,
		Payload:   r.Payload,
		InitInput: r.InitInput,
		From:      r.From,
		To:        r.To,
	}

	radioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        operation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Metadata,
		SenderPrivateKey: r.PrivateKey,
	}
	// 广播交易
	radioTransactionResponse := ts.RadioTransaction(radioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionPrivateContractCreateResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFTransactionPrivateContractCreateResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionPrivateContractCreateResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

// Deprecated
func (ts *TransactionService) PrivateContractCall(r request.BIFTransactionPrivateContractCallRequest) response.BIFTransactionPrivateContractCallResponse {

	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFTransactionPrivateContractCallResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFTransactionPrivateContractCallResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.Type < 0 {
		return response.BIFTransactionPrivateContractCallResponse{
			BIFBaseResponse: exception.INVALID_CONTRACT_TYPE_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	operation := request.BIFPrivateContractCallOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.PRIVATE_CONTRACT_CREATE,
		},
		Type:        common.PRIVATE_CONTRACT_CALL,
		Input:       r.Input,
		DestAddress: r.DestAddress,
		From:        r.From,
		To:          r.To,
	}

	radioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        operation,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Metadata,
		SenderPrivateKey: r.PrivateKey,
	}
	// 广播交易
	radioTransactionResponse := ts.RadioTransaction(radioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionPrivateContractCallResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFTransactionPrivateContractCallResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionPrivateContractCallResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

// EvaluateFee 交易的费用评估
func (ts *TransactionService) EvaluateFee(r request.BIFTransactionEvaluateFeeRequest) response.BIFTransactionEvaluateFeeResponse {

	if ts.url == "" {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if !key.IsAddressValid(r.SourceAddress) {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	// 一、获取交易发起的账号nonce值
	accountGetInfoUrl := common.AccountGetInfoURL(ts.url, r.SourceAddress)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var nonceResponse response.BIFAccountGetNonceResponse
	err = json.Unmarshal(dataByte, &nonceResponse)
	if err != nil {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	if nonceResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: nonceResponse.BIFBaseResponse,
		}
	}
	nonce := nonceResponse.Result.Nonce
	if nonce < common.INIT_ONE {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.INVALID_NONCE_ERROR,
		}
	}
	if r.SignatureNumber < common.INIT_ONE {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.INVALID_SIGNATURENUMBER_ERROR,
		}
	}

	// 二、构建操作、序列化交易
	operationService := GetOperationInstance(ts.url)
	operations, bifBaseResponse := operationService.GetOperations(r.Operation, r.SourceAddress)
	if bifBaseResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: bifBaseResponse,
		}
	}
	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.GasPrice < common.INIT_ZERO {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.INVALID_GAS_AMOUNT_ERROR,
		}
	}
	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}
	dataTemp := hex.EncodeToString([]byte(r.Remarks))
	transaction := proto.EvaluateFeeTransaction{
		SourceAddress: r.SourceAddress,
		Nonce:         nonce + 1,
		FeeLimit:      r.FeeLimit,
		GasPrice:      r.GasPrice,
		Metadata:      dataTemp,
		Operations:    operations,
	}

	transactionItem := make(map[string]proto.EvaluateFeeTransaction)
	transactionItem["transaction_json"] = transaction
	testTransactionRequest := make(map[string]interface{})
	transactionItems := make([]map[string]proto.EvaluateFeeTransaction, 0)
	transactionItems = append(transactionItems, transactionItem)
	testTransactionRequest["items"] = transactionItems
	requestByte, err := json.Marshal(testTransactionRequest)
	if err != nil {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	evaluationFeeUrl := common.TransactionEvaluationFee(ts.url)
	data, err := http.HttpPost(evaluationFeeUrl, requestByte)
	if err != nil {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFTransactionEvaluateFeeResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return response.BIFTransactionEvaluateFeeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (ts *TransactionService) GetTxCacheSize() response.BIFTransactionGetTxCacheSizeResponse {
	getTxCacheSizeUrl := common.GetTxCacheSize(ts.url)
	dataByte, err := http.HttpGet(getTxCacheSizeUrl)
	if err != nil {
		return response.BIFTransactionGetTxCacheSizeResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}
	var getTxCacheSizeResponse response.TransactionGetTxCacheSizeResponse
	err = json.Unmarshal(dataByte, &getTxCacheSizeResponse)
	if err != nil {
		return response.BIFTransactionGetTxCacheSizeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return response.BIFTransactionGetTxCacheSizeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionGetTxCacheSizeResult{
			QueueSize: getTxCacheSizeResponse.QueueSize,
		},
	}
}

func (ts *TransactionService) GetTxCacheData(r request.BIFTransactionCacheRequest) response.BIFTransactionCacheResponse {
	if ts.url == "" {
		return response.BIFTransactionCacheResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.Hash != "" && len(r.Hash) != common.HASH_HEX_LENGTH {
		return response.BIFTransactionCacheResponse{
			BIFBaseResponse: exception.INVALID_HASH_ERROR,
		}
	}
	getTxCacheSizeUrl := common.GetTxCacheData(ts.url, r.Hash)
	dataByte, err := http.HttpGet(getTxCacheSizeUrl)
	if err != nil {
		return response.BIFTransactionCacheResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}
	var res response.BIFTransactionCacheResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFTransactionCacheResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (ts *TransactionService) BatchGasSend(r request.BIFBatchGasSendRequest) response.BIFTransactionGasSendResponse {
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if len(r.Operations) > 100 || len(r.Operations) == 0 {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.OPERATIONS_INVALID_ERROR,
		}
	}
	if r.PrivateKey == "" {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}
	if r.FeeLimit == 0 {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.FeeLimit < common.INIT_ZERO {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}
	if r.GasPrice < common.INIT_ZERO {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: exception.INVALID_GASPRICE_ERROR,
		}
	}
	var operations []request.BIFGasSendOperation
	for _, v := range r.Operations {
		if !key.IsAddressValid(v.DestAddress) {
			return response.BIFTransactionGasSendResponse{
				BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
			}
		}
		if v.Amount == 0 || v.Amount < common.INIT_ZERO {
			return response.BIFTransactionGasSendResponse{
				BIFBaseResponse: exception.INVALID_AMOUNT_ERROR,
			}
		}
		operations = append(operations, request.BIFGasSendOperation{
			DestAddress: v.DestAddress,
			Amount:      v.Amount,
			BIFBaseOperation: request.BIFBaseOperation{
				OperationType: common.GAS_SEND,
			},
		})
	}

	radioTransactionRequest := request.BIFRadioTransactionRequest{
		SenderAddress:    r.SenderAddress,
		FeeLimit:         r.FeeLimit,
		GasPrice:         r.GasPrice,
		Operation:        operations,
		CeilLedgerSeq:    r.CeilLedgerSeq,
		Remarks:          r.Remarks,
		SenderPrivateKey: r.PrivateKey,
	}
	// 广播交易
	radioTransactionResponse := ts.RadioTransaction(radioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionGasSendResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFTransactionGasSendResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionGasSendResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

func (ts *TransactionService) ParseBlob(blob string) response.BIFTransactionParseBlobResponse {
	if blob == "" {
		return response.BIFTransactionParseBlobResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	blobByte, err := hex.DecodeString(blob)
	if err != nil {
		return response.BIFTransactionParseBlobResponse{
			BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
		}
	}
	var transaction proto.Transaction
	err = protobuf.Unmarshal(blobByte, &transaction)
	if err != nil {
		return response.BIFTransactionParseBlobResponse{
			BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
		}
	}

	var operations []response.BIFOperationFormat
	var operation response.BIFOperationFormat
	for _, v := range transaction.Operations {
		var createAccount response.BIFAccountActiviateInfo
		createAccountByte, err := json.Marshal(v.CreateAccount)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}
		err = json.Unmarshal(createAccountByte, &createAccount)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}

		var sendGas response.BIFGasSendInfo
		sendGasByte, err := json.Marshal(v.PayCoin)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}
		err = json.Unmarshal(sendGasByte, &sendGas)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}

		var setMetadata response.BIFAccountSetMetadataInfo
		setMetadataByte, err := json.Marshal(v.SetMetadata)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}
		err = json.Unmarshal(setMetadataByte, &setMetadata)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}

		var setPrivilege response.BIFAccountSetPrivilegeInfo
		setPrivilegeByte, err := json.Marshal(v.SetPrivilege)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}
		err = json.Unmarshal(setPrivilegeByte, &setPrivilege)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}

		var log response.BIFLogInfo
		logByte, err := json.Marshal(v.Log)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}
		err = json.Unmarshal(logByte, &log)
		if err != nil {
			return response.BIFTransactionParseBlobResponse{
				BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
			}
		}

		operation = response.BIFOperationFormat{
			Type:          common.GetOperationType(int(v.Type)),
			SourceAddress: v.SourceAddress,
			Metadata:      string(v.Metadata),
			CreateAccount: createAccount,
			SendGas:       sendGas,
			SetMetadata:   setMetadata,
			SetPrivilege:  setPrivilege,
			Log:           log,
		}
		operations = append(operations, operation)
	}

	return response.BIFTransactionParseBlobResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionParseBlobResult{
			SourceAddress: transaction.SourceAddress,
			FeeLimit:      transaction.FeeLimit,
			GasPrice:      transaction.GasPrice,
			Nonce:         transaction.Nonce,
			Operations:    operations,
			ChainId:       transaction.ChainId,
			Remarks:       string(transaction.Metadata),
		},
	}
}

func (ts *TransactionService) GetBidByHash(hash string) response.BIFTransactionGetBidResponse {
	if ts.url == "" {
		return response.BIFTransactionGetBidResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if hash == "" {
		return response.BIFTransactionGetBidResponse{
			BIFBaseResponse: exception.INVALID_HASH_ERROR,
		}
	}
	r := request.BIFTransactionGetInfoRequest{
		Hash: hash,
	}
	res := ts.GetTransactionInfo(r)
	if res.ErrorCode != 0 {
		return response.BIFTransactionGetBidResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	var bids []string

	for _, v := range res.Result.Transactions {
		for _, op := range v.Transaction.Operations {
			if op.SendGas.DestAddress == common.DDO_CONTRACT {
				var getBidByHashInput response.GetBidByHashInput
				err := json.Unmarshal([]byte(op.SendGas.Input), &getBidByHashInput)
				if err != nil {
					return response.BIFTransactionGetBidResponse{
						BIFBaseResponse: exception.SYSTEM_ERROR,
					}
				}
				bids = append(bids, getBidByHashInput.Params.Document.Id)
			}
		}
	}
	return response.BIFTransactionGetBidResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionGetBidResult{
			Bids: bids,
		},
	}
}
