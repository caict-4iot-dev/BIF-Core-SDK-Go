package request

// BIFBlockGetTransactionsRequest 查询指定区块高度下的所有交易请求体
type BIFBlockGetTransactionsRequest struct {
	BlockNumber int64 `json:"block_number"` // 必填，最新的区块高度，对应底层字段seq
}

// BIFBlockGetInfoRequest 获取指定区块信息请求体
type BIFBlockGetInfoRequest struct {
	BlockNumber int64 `json:"block_number"` // 必填，待查询的区块高度
}

// BIFBlockGetValidatorsRequest 获取指定区块中所有验证节点数请求体
type BIFBlockGetValidatorsRequest struct {
	BlockNumber int64 `json:"block_number"` // 必填，待查询的区块高度，必须大于0
}
