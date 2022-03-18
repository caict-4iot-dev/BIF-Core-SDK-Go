package request

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"

// BIFCreateAccountRequest 创建账户请求体
type BIFCreateAccountRequest struct {
	SenderAddress string `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	DestAddress   string `json:"dest_address"`    // 必填，目标账户地址
	PrivateKey    string `json:"private_key"`     // 必填，交易源账户私钥
	InitBalance   int64  `json:"init_balance"`    // 必填，初始化星火令，单位PT，1 星火令 = 10^8 PT
	CeilLedgerSeq int64  `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Remarks       string `json:"remarks"`         // 可选，用户自定义给交易的备注，16进制格式
	FeeLimit      int64  `json:"fee_limit"`
	GasPrice      int64  `json:"gas_price"`
}

// BIFAccountGetInfoRequest 获取指定的账户信息请求体
type BIFAccountGetInfoRequest struct {
	Address string `json:"address"` // 必填，待查询的区块链账户地址
}

// BIFAccountGetNonceRequest 获取指定账户的nonce值请求体
type BIFAccountGetNonceRequest struct {
	Address string `json:"address"` // 必填，待查询的区块链账户地址
}

// BIFAccountGetBalanceRequest 获取指定账户的星火令的余额请求体
type BIFAccountGetBalanceRequest struct {
	Address string `json:"address"` // 必填，待查询的区块链账户地址
}

// BIFAccountSetMetadatasRequest 设置metadatas请求体
type BIFAccountSetMetadatasRequest struct {
	SenderAddress string `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	PrivateKey    string `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq int64  `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Remarks       string `json:"remarks"`         // 可选，用户自定义给交易的备注，16进制格式
	Key           string `json:"key"`             // 必填，metadata的关键词，长度限制[1, 1024]
	Value         string `json:"value"`           // 必填，metadata的内容，长度限制[0, 256000]
	Version       int64  `json:"version"`         // 选填，metadata的版本
	DeleteFlag    bool   `json:"delete_flag"`     // 选填，是否删除metadata
	FeeLimit      int64  `json:"fee_limit"`
	GasPrice      int64  `json:"gas_price"`
}

// BIFAccountGetMetadatasRequest 获取指定账户的metadatas信息请求体
type BIFAccountGetMetadatasRequest struct {
	Address string `json:"address"` // 必填，待查询的账户地址
	Key     string `json:"key"`     // 选填，metadata关键字，长度限制[1, 1024]，有值为精确查找，无值为全部查找
}

// BIFAccountSetPrivilegeRequest 设置权限请求体
type BIFAccountSetPrivilegeRequest struct {
	SenderAddress  string                      `json:"sender_address"`  // 必填，交易源账号，即交易的发起方
	PrivateKey     string                      `json:"private_key"`     // 必填，交易源账户私钥
	CeilLedgerSeq  int64                       `json:"ceil_ledger_seq"` // 可选，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效
	Remarks        string                      `json:"remarks"`         // 可选，用户自定义给交易的备注，16进制格式
	Signers        []response.BIFSigner        `json:"signers"`         // 选填，签名者权重列表
	TxThreshold    string                      `json:"tx_threshold"`    // 选填，交易门限，大小限制[0, Long.MAX_VALUE]
	TypeThresholds []response.BIFTypeThreshold `json:"type_threshold"`  // 选填，指定类型交易门限
	MasterWeight   string                      // 选填
	FeeLimit       int64                       `json:"fee_limit"`
	GasPrice       int64                       `json:"gas_price"`
}

// Signer 签名者权重列表
type Signer struct {
	Address string `json:"address"` // 签名者区块链账户地址
	Weight  int64  `json:"weight"`  // 选填，metadata的版本
}

// TypeThreshold 指定类型交易门限
type TypeThreshold struct {
	Type      int64 `json:"type"`      // 操作类型，必须大于0
	Threshold int64 `json:"threshold"` // 门限值，大小限制[0, Long.MAX_VALUE]
}

// BIFAccountPrivRequest  获取账户权限请求体
type BIFAccountPrivRequest struct {
	Address string `json:"address"` // 必填，待查询的区块链账户地址
}
