package request

// BIFTransactionGasSendRequest 发送交易请求体
type BIFTransactionGasSendRequest struct {
	SenderAddress string `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	PrivateKey    string `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64  `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	DestAddress   string `json:"dest_address"`    // 必填，发起方地址
	Amount        int64  `json:"amount"`          // 必填，转账金额
	Remarks       string `json:"remarks"`         // 可选，用户自定义给交易的备注，16进制格式
	FeeLimit      int64  `json:"fee_limit"`
	GasPrice      int64  `json:"gas_price"`
}

// BIFTransactionPrivateContractCreateRequest 私有化交易-合约创建请求体
type BIFTransactionPrivateContractCreateRequest struct {
	SenderAddress string   `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	PrivateKey    string   `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64    `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Metadata      string   `json:"metadata"`        // 可选，用户自定义给交易的备注，16进制格式
	Type          int      `json:"type"`            // 必填，合约的语种
	Payload       string   `json:"payload"`         // 必填，对应语种的合约代码
	InitInput     string   `json:"init_input"`      // 必填，合约代码中init方法的入参
	From          string   `json:"from"`            // 必填，发起方加密机公钥
	To            []string `json:"to"`              // 必填，接收方加密机公钥
	FeeLimit      int64    `json:"fee_limit"`
	GasPrice      int64    `json:"gas_price"`
}

// BIFTransactionPrivateContractCallRequest 私有化交易-合约调用请求体
type BIFTransactionPrivateContractCallRequest struct {
	SenderAddress string   `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	PrivateKey    string   `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64    `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Metadata      string   `json:"metadata"`        // 可选，用户自定义给交易的备注，16进制格式
	DestAddress   string   `json:"dest_address"`    // 必填，发起方地址
	Type          int      `json:"type"`            // 必填，合约的语种
	Input         string   `json:"input"`           // 必填，合约代码中init方法的入参
	From          string   `json:"from"`            // 必填，发起方加密机公钥
	To            []string `json:"to"`              // 必填，接收方加密机公钥
	FeeLimit      int64    `json:"fee_limit"`
	GasPrice      int64    `json:"gas_price"`
}

// BIFTransactionGetInfoRequest 根据交易hash查询交易请求体
type BIFTransactionGetInfoRequest struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFRadioTransactionRequest 广播交易请求体
type BIFRadioTransactionRequest struct {
	SenderAddress    string      `json:"sender_address"`
	FeeLimit         int64       `json:"fee_limit"`
	GasPrice         int64       `json:"gas_price"`
	Operation        interface{} `json:"operation"`
	CeilLedgerSeq    int64       `json:"ceil_ledger_seq"`
	Remarks          string      `json:"remarks"`
	SenderPrivateKey string      `json:"privateKey"`
}

// BIFPrivateTransactionSendRequest ...
type BIFPrivateTransactionSendRequest struct {
	From        string
	Payload     string
	To          []string
	DestAddress string
}

type BIFTransactionSubmitRequest struct {
	Serialization string `json:"serialization"`
	SignData      string `json:"sign_data"`
	PublicKey     string `json:"public_key"`
}

// TransactionSubmitRequest 交易提交请求体
type TransactionSubmitRequest struct {
	Items []TransactionSubmit `json:"items"`
}

// TransactionSubmit ...
type TransactionSubmit struct {
	TransactionBlob string      `json:"transaction_blob"`
	Signatures      []Signature `json:"signatures"`
}

// Signature ...
type Signature struct {
	PublicKey string `json:"public_key"`
	SignData  string `json:"sign_data"`
}

// BIFTransactionSerializeRequest ...
type BIFTransactionSerializeRequest struct {
	SourceAddress string      `json:"source_address"`
	Nonce         int64       `json:"nonce"`
	GasPrice      int64       `json:"gas_price"`
	FeeLimit      int64       `json:"fee_limit"`
	Operation     interface{} `json:"operation"`
	CeilLedgerSeq int64       `json:"ceil_ledger_seq"`
	Metadata      string      `json:"metadata"`
}

type BIFTransactionCacheRequest struct {
	Hash string `json:"hash"`
}

// BIFBatchGasSendRequest 批量交易请求体
type BIFBatchGasSendRequest struct {
	SenderAddress string                `json:"sender_address"`
	PrivateKey    string                `json:"private_key"`
	CeilLedgerSeq int64                 `json:"ceil_ledger_seq"`
	Remarks       string                `json:"remarks"`
	FeeLimit      int64                 `json:"fee_limit"`
	GasPrice      int64                 `json:"gas_price"`
	Operations    []BIFGasSendOperation `json:"operations"`
}
