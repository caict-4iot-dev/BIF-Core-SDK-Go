package contract

import (
	"encoding/json"
	"fmt"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"testing"
)

// SDK_INSTANCE_URL 链访问地址
const SDK_INSTANCE_URL = "http://test-bif-core.xinghuo.space"

func TestGetContractInfo(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractGetInfoRequest
	r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
	res := bs.GetContractInfo(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGetContractAddress(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractGetAddressRequest
	r.Hash = "ff6a9d1a0c0011fbb9f51cfb99e4cd5e7c31380046fda3fd6e0daae44d1d4648"
	res := bs.GetContractAddress(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestCheckContractAddress(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractCheckValidRequest
	r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
	res := bs.CheckContractAddress(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestContractQuery(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractCallRequest
	r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
	res := bs.ContractQuery(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestContractInvoke(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractInvokeRequest
	senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	contractAddress := "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
	senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"

	r.SenderAddress = senderAddress
	r.PrivateKey = senderPrivateKey
	r.ContractAddress = contractAddress
	r.BIFAmount = 1
	r.Metadata = "contract invoke"

	res := bs.ContractInvoke(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestContractCreate(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFContractCreateRequest
	senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
	payload := "\"use strict\"; function init(bar){return;} function main(input){let para = JSON.parse(input);if(para.do_foo){let x = {'hello' : 'world'};}} function query(input){return input;}"
	r.SenderAddress = senderAddress
	r.PrivateKey = senderPrivateKey
	r.Metadata = "create contract"
	r.Payload = payload
	r.InitBalance = 1
	r.Type = 0
	r.InitBalance = 1
	r.FeeLimit = 10000000000

	res := bs.ContractCreate(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestBatchContractInvoke(t *testing.T) {
	bs := GetContractInstance(SDK_INSTANCE_URL)
	var r request.BIFBatchContractInvokeRequest
	senderAddress := "did:bid:ef7zyvBtyg22NC4qDHwehMJxeqw6Mmrh"
	contractAddress := "did:bid:eftzENB3YsWymQnvsLyF4T2ENzjgEg41"
	senderPrivateKey := "priSPKr2dgZTCNj1mGkDYyhyZbCQhEzjQm7aEAnfVaqGmXsW2x"

	r.SenderAddress = senderAddress
	r.PrivateKey = senderPrivateKey

	keyPair01, err := key.GetBidAndKeyPair()
	if err != nil {
		t.Error(err)
	}
	destAddress01 := keyPair01.GetEncAddress()
	keyPair02, err := key.GetBidAndKeyPair()
	if err != nil {
		t.Error(err)
	}
	destAddress02 := keyPair02.GetEncAddress()

	input01 := "{\"method\":\"creation\",\"params\":{\"document\":{\"@context\": [\"https://w3.org/ns/did/v1\"],\"context\": \"https://w3id.org/did/v1\"," +
		"\"id\": \"" + destAddress01 + "\", \"version\": \"1\"}}}"
	input02 := "{\"method\":\"creation\",\"params\":{\"document\":{\"@context\": [\"https://w3.org/ns/did/v1\"],\"context\": \"https://w3id.org/did/v1\"," +
		"\"id\": \"" + destAddress02 + "\", \"version\": \"1\"}}}"

	var operations []request.BIFContractInvokeOperation
	var operation01 request.BIFContractInvokeOperation
	var operation02 request.BIFContractInvokeOperation
	operation01.ContractAddress = contractAddress
	operation01.Input = input01
	operation02.ContractAddress = contractAddress
	operation02.Input = input02
	operations = append(operations, operation01, operation02)
	r.Operations = operations
	res := bs.BatchContractInvoke(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}
