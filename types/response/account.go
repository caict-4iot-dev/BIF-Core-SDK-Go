package response

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/proto"

// BIFAccountGetInfoResponse 获取指定的账户信息返回体
type BIFAccountGetInfoResponse struct {
	BIFBaseResponse
	Result BIFAccountGetInfoResult `json:"result"`
}

type BIFAccountGetInfoResult struct {
	Address string `json:"address"` // 账户地址
	Balance int64  `json:"balance"` // 账户余额，单位PT，1 星火令 = 10^8 PT, 必须大于0
	Nonce   int64  `json:"nonce"`   // 账户交易序列号，必须大于0
}

// BIFAccountGetBalanceResponse 获取指定账户的星火令的余额返回体
type BIFAccountGetBalanceResponse struct {
	BIFBaseResponse
	Result BIFAccountGetBalanceResult `json:"result"`
}

type BIFAccountGetBalanceResult struct {
	Balance int64 `json:"balance"` // 余额
}

// BIFAccountSetMetadatasResponse 设置metadatas返回体
type BIFAccountSetMetadatasResponse struct {
	BIFBaseResponse
	Result BIFAccountSetMetadataResult `json:"result"`
}

type BIFAccountSetMetadataResult struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFAccountGetMetadatasResponse 获取指定账户的metadata信息返回体
type BIFAccountGetMetadatasResponse struct {
	BIFBaseResponse
	Result BIFAccountGetMetadatasResult `json:"result"`
}

type BIFAccountGetMetadatasResult struct {
	Metadatas []BIFMetadataInfo `json:"metadatas"`
}

// BIFMetadataInfo 账户信息
type BIFMetadataInfo struct {
	Key     string `json:"key"`     // metadata的关键词
	Value   string `json:"value"`   // metadata的内容
	Version int64  `json:"version"` // metadata的版本
}

// BIFAccountSetPrivilegeResponse 设置权限返回体
type BIFAccountSetPrivilegeResponse struct {
	BIFBaseResponse
	Result BIFAccountSetPrivilegeResult `json:"result"`
}

type BIFAccountSetPrivilegeResult struct {
	Hash string `json:"hash"` // 交易hash
}

// BIFAccountPrivResponse 获取账户权限返回体
type BIFAccountPrivResponse struct {
	BIFBaseResponse
	Result BIFAccountPrivResult `json:"result"`
}

type BIFAccountPrivResult struct {
	Address string  `json:"address"` // 账户地址
	Priv    BIFPriv `json:"priv"`    // 账户权限
}

// BIFAccountGetNonceResponse 获取指定账户的nonce值返回体
type BIFAccountGetNonceResponse struct {
	BIFBaseResponse
	Result BIFAccountGetNonceResult `json:"result"`
}

type BIFAccountGetNonceResult struct {
	Nonce int64 `json:"nonce"`
}

type BIFCreateAccountResponse struct {
	BIFBaseResponse
	Result BIFAccountCreateAccountResult `json:"result"`
}

type BIFAccountCreateAccountResult struct {
	Hash string `json:"hash"`
}

type BIFAccountSetPrivilegeOperationResponse struct {
	BIFBaseResponse
	Result BIFAccountSetPrivilegeOperationResult `json:"result"`
}

type BIFAccountSetPrivilegeOperationResult struct {
	Operation *proto.Operation
}
