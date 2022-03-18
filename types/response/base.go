package response

// BIFBaseResponse SDK响应返回结构体
type BIFBaseResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}
