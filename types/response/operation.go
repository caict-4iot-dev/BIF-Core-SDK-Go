package response

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/proto"

type BIFBaseOperationResult struct {
	Operation *proto.Operation `json:"operation"`
}

type AccountActivateResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFGasSendOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFPrivateContractCallOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFContractCreateOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFContractInvokeOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFPrivateContractCreateOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type PrivateContractCallOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}

type BIFAccountSetMetadataOperationResponse struct {
	BIFBaseResponse
	Result BIFBaseOperationResult `json:"result"`
}
