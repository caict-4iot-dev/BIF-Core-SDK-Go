package response

// BIFBlockGetNumberResponse ...
type BIFBlockGetNumberResponse struct {
	BIFBaseResponse
	Result BIFBlockGetNumberResult `json:"result"`
}

type BIFBlockGetNumberResult struct {
	Header BIFBlockNumber `json:"header"` // 区块头
}

// BlockHeader 区块头
type BIFBlockNumber struct {
	BlockNumber int64 `json:"seq"` // 最新的区块高度，对应底层字段seq
}

// BIFBlockGetTransactionsResponse 查询指定区块高度下的所有交易响应体
type BIFBlockGetTransactionsResponse struct {
	BIFBaseResponse
	Result BIFBlockGetTransactionsResult `json:"result"`
}

type BIFBlockGetTransactionsResult struct {
	TotalCount   int64                   `json:"total_count"`  // 返回的总交易数
	Transactions []BIFTransactionHistory `json:"transactions"` // 交易内容
}

type BIFSignature struct {
	SignData  string `json:"sign_data"`
	PublicKey string `json:"public_key"`
}

type BIFTransactionInfo struct {
	SourceAddress string         `json:"source_address"`
	FeeLimit      int64          `json:"fee_limit"`
	GasPrice      int64          `json:"gas_price"`
	Nonce         int64          `json:"nonce"`
	Metadata      string         `json:"metadata"`
	Operations    []BIFOperation `json:"operations"`
	ChainId       int64          `json:"chain_id"`
}

type BIFOperation struct {
	Type          int                        `json:"type"`
	SourceAddress string                     `json:"source_address"`
	Metadata      string                     `json:"metadata"`
	CreateAccount BIFAccountActiviateInfo    `json:"create_account"`
	SendGas       BIFGasSendInfo             `json:"pay_coin"` // pay_coin
	SetMetadata   BIFAccountSetMetadataInfo  `json:"set_metadata"`
	SetPrivilege  BIFAccountSetPrivilegeInfo `json:"set_privilege"`
	Log           BIFLogInfo                 `json:"log"`
}

type BIFLogInfo struct {
	Topic string   `json:"topic"`
	Datas []string `json:"datas"`
}

type BIFAccountSetPrivilegeInfo struct {
	MasterWeight   string             `json:"master_weight"`
	Signers        []BIFSigner        `json:"signers"`
	TxThreshold    string             `json:"tx_threshold"`
	TypeThresholds []BIFTypeThreshold `json:"type_thresholds"`
}

type BIFAccountSetMetadataInfo struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Version    int64  `json:"version"`
	DeleteFlag bool   `json:"delete_flag"`
}

type BIFGasSendInfo struct {
	DestAddress string `json:"dest_address"`
	Amount      int64  `json:"amount"`
	Input       string `json:"input"`
}

type BIFAccountActiviateInfo struct {
	DestAddress string            `json:"dest_address"`
	Contract    BIFContractInfo   `json:"contract"`
	Priv        BIFPriv           `json:"priv"`
	Metadatas   []BIFMetadataInfo `json:"metadatas"`
	InitBalance int64             `json:"init_balance"`
	InitInput   string            `json:"init_input"`
}

type BIFPriv struct {
	MasterWeight int64        `json:"master_weight"`
	Signers      []BIFSigner  `json:"signers"`
	Thresholds   BIFThreshold `json:"thresholds"` // thresholds
}

type BIFThreshold struct {
	TxThreshold    int64              `json:"tx_threshold"`
	TypeThresholds []BIFTypeThreshold `json:"type_thresholds"`
}

type BIFTypeThreshold struct {
	Type      int   `json:"type"`
	Threshold int64 `json:"threshold"`
}

type BIFSigner struct {
	Address string `json:"address"`
	Weight  int64  `json:"weight"`
}

// BIFBlockGetInfoResponse 获取指定区块信息响应体
type BIFBlockGetInfoResponse struct {
	BIFBaseResponse
	Result BIFBlockGetInfoResult `json:"result"`
}

type BIFBlockGetInfoResult struct {
	Header BIFBlockHeader `json:"header"` // 区块信息
}

// BIFBlockHeader 区块信息
type BIFBlockHeader struct {
	ConfirmTime int64 `json:"close_time"` // 区块确认时间
	Number      int64 `json:"seq"`        // 区块高度
	TxCount     int64 `json:"tx_count"`   // 交易总量
	Version     int64 `json:"version"`    // 区块版本
}

// BIFBlockGetLatestInfoResponse 获取最新区块信息响应体
type BIFBlockGetLatestInfoResponse struct {
	BIFBaseResponse
	Result BIFBlockGetLatestInfoResult `json:"result"`
}

type BIFBlockGetLatestInfoResult struct {
	Header BIFBlockHeader `json:"header"` // 区块信息
}

// BIFBlockGetValidatorsResponse 获取指定区块中所有验证节点数响应体
type BIFBlockGetValidatorsResponse struct {
	BIFBaseResponse
	Result BIFBlockGetValidatorsResult `json:"result"`
}

type BIFBlockGetValidatorsResult struct {
	Validators []string `json:"validators"` // 验证节点列表
}

// BIFBlockGetLatestValidatorsResponse 获取最新区块中所有验证节点数响应体
type BIFBlockGetLatestValidatorsResponse struct {
	BIFBaseResponse
	Result BIFBlockGetValidatorsResult `json:"result"`
}

type BIFBlockGetLatestValidatorsResult struct {
	Validators []string `json:"validators"` // 验证节点列表
}
