package response

// BIFContractCheckValidResponse 检测合约账户的有效性返回体
type BIFContractCheckValidResponse struct {
	BIFBaseResponse
	Result BIFContractCheckValidResult `json:"result"`
}

type BIFContractCheckValidResult struct {
	IsValid bool `json:"is_valid"` // 是否有效
}

// BIFContractCreateResponse 创建合约响应体
type BIFContractCreateResponse struct {
	BIFBaseResponse
	Result BIFContractInvokeResult
}

// BIFContractGetInfoResponse 查询合约代码响应体
type BIFContractGetInfoResponse struct {
	BIFBaseResponse
	Result BIFContractGetInfoResult `json:"result"`
}

type BIFContractGetInfoResult struct {
	Contract BIFContractInfo `json:"contract"` // 合约信息
}

// BIFContractInfo 合约信息
type BIFContractInfo struct {
	Type    int    `json:"type"`    // 合约类型，默认0
	Payload string `json:"payload"` // 合约代码
}

// BIFContractGetAddressResponse 根据交易Hash查询合约地址响应体
type BIFContractGetAddressResponse struct {
	BIFBaseResponse
	Result BIFContractGetAddressResult `json:"result"`
}

type BIFContractGetAddressResult struct {
	ContractAddressInfos []ContractAddressInfo `json:"contract_address_infos"` // 合约地址列表
}

// ContractAddressInfo 合约信息
type ContractAddressInfo struct {
	ContractAddress string `json:"contract_address"` // 合约地址
	OperationIndex  int    `json:"operation_index"`  // 所在操作的下标
}

// BIFContractCallResponse 合约查询接口响应体
type BIFContractCallResponse struct {
	BIFBaseResponse
	Result BIFContractCallResult `json:"result"`
}

type BIFContractCallResult struct {
	QueryRets []QueryRetsResult `json:"query_rets"` // 查询结果集
}

type QueryRetsResult struct {
	Result QueryResult `json:"result"`
}

type QueryResult struct {
	//Type  string `json:"type"`
	//Value string `json:"value"`
	Data    string `json:"data"`
	Gasused int64  `json:"gasused"`
}

// BIFContractInvokeResponse 合约调用响应体
type BIFContractInvokeResponse struct {
	BIFBaseResponse
	Result BIFContractInvokeResult `json:"result"`
}

type BIFContractInvokeResult struct {
	Hash string `json:"hash"` // 交易hash
}
