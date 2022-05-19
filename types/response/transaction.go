package response

// BIFTransactionGasSendResponse 发送交易响应体
type BIFTransactionGasSendResponse struct {
	BIFBaseResponse
	Result BIFTransactionGasSendResult `json:"result"`
}

type BIFTransactionGasSendResult struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFTransactionPrivateContractCreateResponse 私有化交易-合约创建响应体
type BIFTransactionPrivateContractCreateResponse struct {
	BIFBaseResponse
	Result BIFTransactionPrivateContractCreateResult `json:"result"`
}

type BIFTransactionPrivateContractCreateResult struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFTransactionPrivateContractCallResponse 私有化交易-合约调用响应体
type BIFTransactionPrivateContractCallResponse struct {
	BIFBaseResponse
	Result BIFTransactionPrivateContractCallResult `json:"result"`
}

type BIFTransactionPrivateContractCallResult struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFTransactionGetInfoResponse 根据交易hash查询交易响应体
type BIFTransactionGetInfoResponse struct {
	BIFBaseResponse
	Result BIFTransactionGetInfoResult `json:"result"`
}

type BIFTransactionGetInfoResult struct {
	TotalCount   int64                   `json:"total_count"`  // 返回的总交易数
	Transactions []BIFTransactionHistory `json:"transactions"` // 交易内容
}

// BIFTransactionHistory 交易内容
type BIFTransactionHistory struct {
	Fee              int64              `json:"actual_fee"` // 交易实际费用
	ConfirmTime      int64              `json:"close_time"` // 交易确认时间
	ErrorCode        int64              `json:"error_code"` // 交易错误码
	ErrorDesc        string             `json:"error_desc"` // 交易描述
	Hash             string             `json:"hash"`       // 交易hash
	LedgerSeq        int64              `json:"ledger_seq"` // 区块序列号
	TxSize           int64              `json:"tx_size"`
	Signatures       []BIFSignature     `json:"signatures"`
	ContractTxHashes []string           `json:"contract_tx_hashes"`
	Transaction      BIFTransactionInfo `json:"transaction"` // 交易内容列表
}

// TransactionInfo 交易内容列表
type TransactionInfo struct {
	Signatures Signature `json:"signatures"` // 签名列表

}

// Signature 签名信息
type Signature struct {
	SignData  int64 `json:"sign_data"`  // 签名后数据
	PublicKey int64 `json:"public_key"` // 公钥
	TxSize    int64 `json:"tx_size"`    // 交易大小
}

type BIFTransactionSubmitResponse struct {
	BIFBaseResponse
	Result BIFTransactionSubmitResult `json:"result"`
}

type BIFTransactionSubmitResult struct {
	Hash string `json:"hash"`
}

type TransactionSubmitResponse struct {
	Results      []TransactionSubmitResult `json:"results"`
	SuccessCount int                       `json:"success_count"`
}

type TransactionSubmitResult struct {
	BIFBaseResponse
	Hash string `json:"hash"`
}

type BIFTransactionSerializeResponse struct {
	BIFBaseResponse
	Result BIFTransactionSerializeResult `json:"result"`
}

type BIFTransactionSerializeResult struct {
	TransactionBlob string `json:"transaction_blob"`
	Hash            string `json:"hash"`
}

type BIFRadioTransactionResponse struct {
	BIFBaseResponse
	Result BIFRadioTransactionResult `json:"result"`
}

type BIFRadioTransactionResult struct {
	Hash string `json:"hash"`
}

type BIFPrivateTransactionSendResponse struct {
	BIFBaseResponse
	PriTxHash string `json:"pri_tx_hash"`
}

type BIFTransactionEvaluateFeeResponse struct {
	BIFBaseResponse
	Result BIFTransactionEvaluateFeeResult `json:"result"`
}

type BIFTransactionEvaluateFeeResult struct {
	Txs []BIFTestTx `json:"txs"`
}

type BIFTestTx struct {
	TransactionEnv BIFTestTransactionFees `json:"transaction_env"`
}

type BIFTestTransactionFees struct {
	TransactionFees BIFTransactionFees `json:"transaction"`
}

type BIFTransactionFees struct {
	FeeLimit int64 `json:"fee_limit"`
	GasPrice int64 `json:"gas_price"`
}

type BIFTransactionGetTxCacheSizeResponse struct {
	BIFBaseResponse
	Result BIFTransactionGetTxCacheSizeResult `json:"result"`
}

type BIFTransactionGetTxCacheSizeResult struct {
	QueueSize int64 `json:"queue_size"`
}

type TransactionGetTxCacheSizeResponse struct {
	QueueSize int64 `json:"queue_size"`
}
