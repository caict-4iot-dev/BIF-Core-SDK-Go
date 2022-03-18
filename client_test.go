package bif_core_sdk_go

import (
	"fmt"
	"testing"
)

// SDK_INSTANCE_URL 链访问地址
const SDK_INSTANCE_URL = "http://test-bif-core.xinghuo.space"

func TestGetInstance(t *testing.T) {
	_, err := GetInstance(SDK_INSTANCE_URL)
	if err != nil {
		t.Fatal("get sdk instance is failed, err:", err)
	}
}

func TestGetBIFAccountService(t *testing.T) {
	sdk, err := GetInstance(SDK_INSTANCE_URL)
	if err != nil {
		t.Fatal("get sdk instance is failed, err:", err)
	}
	service := sdk.GetBIFAccountService()
	if service == nil {
		t.Error("GetBIFAccountService is failed")
	}
}

func TestGetBlockService(t *testing.T) {
	sdk, err := GetInstance(SDK_INSTANCE_URL)
	if err != nil {
		t.Fatal("get sdk instance is failed, err:", err)
	}
	service := sdk.GetBlockService()
	if service == nil {
		t.Error("GetBlockService is failed")
	}

	// 查询块高度
	res := service.GetBlockNumber()
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	fmt.Printf("result: %+v \n", res.Result)
}

func TestGetTransactionService(t *testing.T) {
	sdk, err := GetInstance(SDK_INSTANCE_URL)
	if err != nil {
		t.Fatal("get sdk instance is failed, err:", err)
	}
	service := sdk.GetTransactionService()
	if service == nil {
		t.Error("GetTransactionService is failed")
	}
}

func TestGetContractService(t *testing.T) {
	sdk, err := GetInstance(SDK_INSTANCE_URL)
	if err != nil {
		t.Fatal("get sdk instance is failed, err:", err)
	}
	service := sdk.GetContractService()
	if service == nil {
		t.Error("GetContractService is failed")
	}
}
