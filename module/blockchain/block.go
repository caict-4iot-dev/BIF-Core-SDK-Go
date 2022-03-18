package blockchain

import (
	"encoding/json"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/exception"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"
)

// BIFBlockService block interface
type BIFBlockService interface {
	// GetBlockNumber 查询最新的区块高度
	GetBlockNumber() response.BIFBlockGetNumberResponse
	// GetTransactions 查询指定区块高度下的所有交易
	GetTransactions(bifBlockGetTransactionsRequest request.BIFBlockGetTransactionsRequest) response.BIFBlockGetTransactionsResponse
	// GetBlockInfo 获取指定区块信息
	GetBlockInfo(bifBlockGetInfoRequest request.BIFBlockGetInfoRequest) response.BIFBlockGetInfoResponse
	// GetBlockLatestInfo 获取最新区块信息
	GetBlockLatestInfo() response.BIFBlockGetLatestInfoResponse
	// GetValidators 获取指定区块中所有验证节点数
	GetValidators(bifBlockGetValidatorsRequest request.BIFBlockGetValidatorsRequest) response.BIFBlockGetValidatorsResponse
	// GetLatestValidators 获取最新区块中所有验证节点数
	GetLatestValidators() response.BIFBlockGetLatestValidatorsResponse
}

// BlockService ...
type BlockService struct {
	url string
}

func GetBlockInstance(url string) *BlockService {
	return &BlockService{
		url,
	}
}

func (bs *BlockService) GetBlockNumber() response.BIFBlockGetNumberResponse {

	if bs.url == "" {
		return response.BIFBlockGetNumberResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	getNumberURL := common.BlockGetNumberURL(bs.url)
	dataByte, err := http.HttpGet(getNumberURL)
	if err != nil {
		return response.BIFBlockGetNumberResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetNumberResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {

		return response.BIFBlockGetNumberResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (bs *BlockService) GetTransactions(r request.BIFBlockGetTransactionsRequest) response.BIFBlockGetTransactionsResponse {

	if bs.url == "" {
		return response.BIFBlockGetTransactionsResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}

	if r.BlockNumber < common.INIT_ONE_L {
		return response.BIFBlockGetTransactionsResponse{
			BIFBaseResponse: exception.INVALID_BLOCKNUMBER_ERROR,
		}
	}
	getTransactionsUrl := common.BlockGetTransactionsURL(bs.url, r.BlockNumber)
	dataByte, err := http.HttpGet(getTransactionsUrl)
	if err != nil {
		return response.BIFBlockGetTransactionsResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetTransactionsResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFBlockGetTransactionsResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (bs *BlockService) GetValidators(r request.BIFBlockGetValidatorsRequest) response.BIFBlockGetValidatorsResponse {

	if bs.url == "" {
		return response.BIFBlockGetValidatorsResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.BlockNumber < common.INIT_ONE_L {
		return response.BIFBlockGetValidatorsResponse{
			BIFBaseResponse: exception.INVALID_BLOCKNUMBER_ERROR,
		}
	}
	getInfoUrl := common.BlockGetValidatorsURL(bs.url, r.BlockNumber)
	dataByte, err := http.HttpGet(getInfoUrl)
	if err != nil {
		return response.BIFBlockGetValidatorsResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetValidatorsResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFBlockGetValidatorsResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (bs *BlockService) GetBlockInfo(r request.BIFBlockGetInfoRequest) response.BIFBlockGetInfoResponse {
	if bs.url == "" {
		return response.BIFBlockGetInfoResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}

	if r.BlockNumber < common.INIT_ONE_L {
		return response.BIFBlockGetInfoResponse{
			BIFBaseResponse: exception.INVALID_BLOCKNUMBER_ERROR,
		}
	}

	getInfoUrl := common.BlockGetInfoURL(bs.url, r.BlockNumber)
	dataByte, err := http.HttpGet(getInfoUrl)
	if err != nil {
		return response.BIFBlockGetInfoResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetInfoResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFBlockGetInfoResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (bs *BlockService) GetBlockLatestInfo() response.BIFBlockGetLatestInfoResponse {
	if bs.url == "" {
		return response.BIFBlockGetLatestInfoResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	getInfoUrl := common.BlockGetLatestInfoURL(bs.url)
	dataByte, err := http.HttpGet(getInfoUrl)
	if err != nil {
		return response.BIFBlockGetLatestInfoResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetLatestInfoResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFBlockGetLatestInfoResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}

func (bs *BlockService) GetLatestValidators() response.BIFBlockGetLatestValidatorsResponse {
	if bs.url == "" {
		return response.BIFBlockGetLatestValidatorsResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	getInfoUrl := common.BlockGetLatestValidatorsURL(bs.url)
	dataByte, err := http.HttpGet(getInfoUrl)
	if err != nil {
		return response.BIFBlockGetLatestValidatorsResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.BIFBlockGetLatestValidatorsResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFBlockGetLatestValidatorsResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	return res
}
