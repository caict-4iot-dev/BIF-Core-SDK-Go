package blockchain

import (
	"fmt"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"testing"
)

// SDK_INSTANCE_URL 链访问地址
const SDK_INSTANCE_URL = "http://test.bifcore.bitfactory.cn"

func TestGetBlockNumber(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	res := bs.GetBlockNumber()
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Println("blockNumber:", res.Result.Header.BlockNumber)
}

func TestGetTransactions(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	var r request.BIFBlockGetTransactionsRequest
	r.BlockNumber = 617247
	res := bs.GetTransactions(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}

func TestGetValidators(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	var r request.BIFBlockGetValidatorsRequest
	r.BlockNumber = 617247
	res := bs.GetValidators(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}

func TestGetBlockInfo(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	var r request.BIFBlockGetInfoRequest
	r.BlockNumber = 617247
	res := bs.GetBlockInfo(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}

func TestGetBlockLatestInfo(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	res := bs.GetBlockLatestInfo()
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}

func TestGetLatestValidators(t *testing.T) {
	bs := GetBlockInstance(SDK_INSTANCE_URL)
	res := bs.GetLatestValidators()
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}
