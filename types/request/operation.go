package request

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"

type BIFBaseOperation struct {
	OperationType int    `json:"operation_type"`
	SourceAddress string `json:"source_address"`
	Metadata      string `json:"metadata"`
}

type BIFAccountActivateOperation struct {
	BIFBaseOperation
	DestAddress string
	InitBalance int64
}

type BIFAccountSetPrivilegeOperation struct {
	BIFBaseOperation
	MasterWeight   string
	Signers        []response.BIFSigner
	TxThreshold    string
	TypeThresholds []response.BIFTypeThreshold
}

type BIFAccountSetMetadataOperation struct {
	BIFBaseOperation
	Key        string
	Value      string
	Version    int64
	DeleteFlag bool
}

type BIFPrivateContractCreateOperation struct {
	BIFBaseOperation
	Type      int
	Payload   string
	InitInput string
	From      string
	To        []string
}

type BIFGasSendOperation struct {
	BIFBaseOperation
	DestAddress string `json:"dest_address"`
	Amount      int64  `json:"amount"`
}

type BIFPrivateContractCallOperation struct {
	BIFBaseOperation
	Type        int
	Input       string
	From        string
	To          []string
	DestAddress string
}

type BIFContractCreateOperation struct {
	BIFBaseOperation
	InitBalance int64
	Type        int
	Payload     string
	InitInput   string
}

type BIFContractInvokeOperation struct {
	BIFBaseOperation
	ContractAddress string
	BifAmount       int64
	Input           string
}

// BIFTransactionEvaluateFeeRequest 交易费用评估请求体
type BIFTransactionEvaluateFeeRequest struct {
	SourceAddress   string      // 模拟交易的原地址
	Nonce           int64       // 在原账号基础上加1
	SignatureNumber int64       // 签名个数，默认为1；不填写系统会设置为1
	Metadata        string      // 可选，签名个数
	Operation       interface{} // 操作列表
}
