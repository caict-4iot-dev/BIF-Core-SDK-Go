package request

// BIFContractCheckValidRequest 检测合约账户的有效性请求体
type BIFContractCheckValidRequest struct {
	ContractAddress string `json:"contract_address"` // 待检测的合约账户地址
}

// BIFContractCreateRequest 创建合约请求体
type BIFContractCreateRequest struct {
	SenderAddress string `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	FeeLimit      int64  `json:"fee_limit"`       // 可选，交易花费的手续费，默认1000000L
	PrivateKey    string `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64  `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Metadata      string `json:"metadata"`        // 可选，用户自定义给交易的备注，16进制格式
	InitBalance   int64  `json:"init_balance"`    // 必填，给合约账户的初始化星火令，单位PT，1 星火令 = 10^8 PT, 大小限制[1, Long.MAX_VALUE]
	Type          int    `json:"type"`            // 选填，合约的类型，默认是0 , 0: javascript，1 :evm 。
	Payload       string `json:"payload"`         // 必填，对应语种的合约代码
	InitInput     string `json:"init_input"`      // 选填，合约代码中init方法的入参
	GasPrice      int64  `json:"gas_price"`
}

// BIFContractGetInfoRequest 查询合约代码请求体
type BIFContractGetInfoRequest struct {
	ContractAddress string `json:"contract_address"` // 待查询的合约账户地址
}

// BIFContractGetAddressRequest 根据交易Hash查询合约地址请求体
type BIFContractGetAddressRequest struct {
	Hash string `json:"hash"` // 创建合约交易的hash
}

// BIFContractCallRequest 合约查询接口请求体
type BIFContractCallRequest struct {
	SourceAddress   string `json:"source_address"`   // 选填，合约触发账户地址
	ContractAddress string `json:"contract_address"` // 必填，合约账户地址
	Input           string `json:"input"`            // 选填，合约入参
	FeeLimit        int64  `json:"fee_limit"`
	GasPrice        int64  `json:"gas_price"`
	Type            int    `json:"type"` // 选填，合约类型 默认是0 , 0: javascript，1 :evm
}

// BIFContractInvokeRequest 合约调用请求体
type BIFContractInvokeRequest struct {
	SenderAddress   string `json:"sender_address"`   // 必填，交易源账号，即交易的发起方
	FeeLimit        int64  `json:"fee_limit"`        // 可选，交易花费的手续费，默认1000000L
	PrivateKey      string `json:"private_key"`      // 必填，交易源账户私钥
	CeilLedgerSeq   int64  `json:"ceil_ledger_seq"`  // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Metadata        string `json:"metadata"`         // 可选，用户自定义给交易的备注，16进制格式
	ContractAddress string `json:"contract_address"` // 必填，合约账户地址
	BIFAmount       int64  `json:"bif_amount"`       // 必填，转账金额
	Input           string `json:"input"`            // 选填，待触发的合约的main()入参
	GasPrice        int64  `json:"gas_price"`
}

type BIFBatchContractInvokeRequest struct {
	SenderAddress string                       `json:"sender_address"` // 必填，交易源账号，即交易的发起方
	FeeLimit      int64                        `json:"fee_limit"`      // 可选，交易花费的手续费，默认1000000L
	Operations    []BIFContractInvokeOperation `json:"operations"`
	PrivateKey    string                       `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64                        `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Remarks       string                       `json:"remarks"`         // 可选，用户自定义给交易的备注，16进制格式
	GasPrice      int64                        `json:"gas_price"`
}
