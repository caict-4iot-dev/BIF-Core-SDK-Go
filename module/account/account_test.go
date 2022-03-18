package account

import (
	"encoding/json"
	"fmt"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"testing"
)

// SDK_INSTANCE_URL 链访问地址
const SDK_INSTANCE_URL = "http://test-bif-core.xinghuo.space"

func TestCreateAccount(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
	destAddress := "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
	r := request.BIFCreateAccountRequest{
		SenderAddress: senderAddress,
		DestAddress:   destAddress,
		PrivateKey:    senderPrivateKey,
		InitBalance:   1000000,
		Remarks:       "init account",
	}

	res := as.CreateAccount(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetAccount(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
	r := request.BIFAccountGetInfoRequest{
		Address: accountAddress,
	}
	res := as.GetAccount(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetNonce(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
	r := request.BIFAccountGetNonceRequest{
		Address: accountAddress,
	}
	res := as.GetNonce(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetAccountBalance(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
	r := request.BIFAccountGetBalanceRequest{
		Address: accountAddress,
	}
	res := as.GetAccountBalance(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestSetMetadata(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	r := request.BIFAccountSetMetadatasRequest{
		SenderAddress: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
		PrivateKey:    "priSPKe86UJsnJ3WTDtLViP5ii8WTZKCXRMJmmqkDBWHq1eyMy",
		Remarks:       "set remarks",
		Key:           "20220101-01",
		Value:         "metadata-20220101-01",
	}
	res := as.SetMetadatas(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetAccountMetadata(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	r := request.BIFAccountGetMetadatasRequest{
		Address: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
		Key:     "20220101-01",
	}
	res := as.GetAccountMetadatas(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestSetPrivilege(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	r := request.BIFAccountSetPrivilegeRequest{
		SenderAddress: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
		PrivateKey:    "priSPKe86UJsnJ3WTDtLViP5ii8WTZKCXRMJmmqkDBWHq1eyMy",
		Remarks:       "set privilege",
		TxThreshold:   "0",
		//MasterWeight:  "1",
	}
	res := as.SetPrivilege(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetAccountPriv(t *testing.T) {

	as := GetAccountInstance(SDK_INSTANCE_URL)
	// 初始化请求参数
	r := request.BIFAccountPrivRequest{
		Address: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
	}
	res := as.GetAccountPriv(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}
