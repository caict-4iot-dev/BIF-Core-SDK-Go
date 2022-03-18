package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"
	"testing"
)

func TestGetTransactionInfo(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	var r request.BIFTransactionGetInfoRequest
	r.Hash = "2c0a445f603bdef7e4cfe5f63650f201cda3315b7c560edb79e3fcef611c5f8e"
	res := ts.GetTransactionInfo(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestGasSend(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	var r request.BIFTransactionGasSendRequest
	//r.SenderAddress = "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	//r.PrivateKey = "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
	//r.DestAddress = "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
	r.SenderAddress = "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
	r.PrivateKey = "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
	r.DestAddress = "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	r.Metadata = "gas send"
	r.Amount = 100000

	res := ts.GasSend(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	// "hash":"8034d5b3f3d0ee73e3a5139fe17d340bdfc896da49b8e6d5dcc4b704c959e306"
	fmt.Println("res: ", string(dataByte))
}

func TestPrivateContractCreate(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	var r request.BIFTransactionPrivateContractCreateRequest
	r.SenderAddress = "did:bid:efnVUgqQFfYeu97ABf6sGm3WFtVXHZB2"
	r.PrivateKey = "priSPKkWVk418PKAS66q4bsiE2c4dKuSSafZvNWyGGp2sJVtXL"
	r.Payload = "\"use strict\";function queryBanance(address)\r\n{return \" test query private contract sdk_3\";}\r\nfunction sendTx(to,amount)\r\n{return Chain.payCoin(to,amount);}\r\nfunction init(input)\r\n{return;}\r\nfunction main(input)\r\n{let args=JSON.parse(input);if(args.method===\"sendTx\"){return sendTx(args.params.address,args.params.amount);}}\r\nfunction query(input)\r\n{let args=JSON.parse(input);if(args.method===\"queryBanance\"){return queryBanance(args.params.address);}}"
	r.From = "sX46dMvKzKgH/SByjBs0uCROD9paCc/tF6WwcgUx3nA="
	r.To = []string{"Pz8tQqi4DZcL5Vrh/GXS20vZ4oqaiNyFxG0B9xAJmhw="}
	r.Metadata = "init account"
	r.Type = 0

	res := ts.PrivateContractCreate(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	// "hash":"278a1f189e235fee846baf22b4ff699e702f9dd407d2361bbfb41159d57f4a2f"
	fmt.Println("res: ", string(dataByte))
}

func TestPrivateContractCall(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	var r request.BIFTransactionPrivateContractCallRequest
	r.SenderAddress = "did:bid:efnVUgqQFfYeu97ABf6sGm3WFtVXHZB2"
	r.PrivateKey = "priSPKkWVk418PKAS66q4bsiE2c4dKuSSafZvNWyGGp2sJVtXL"
	r.Input = "{\"method\":\"queryBanance\",\"params\":{\"address\":\"567890哈哈=======\"}}"
	r.From = "sX46dMvKzKgH/SByjBs0uCROD9paCc/tF6WwcgUx3nA="
	r.To = []string{"Pz8tQqi4DZcL5Vrh/GXS20vZ4oqaiNyFxG0B9xAJmhw="}
	r.DestAddress = "did:bid:efTuswkPE1HP9Uc7vpNbRVokuQqhxaCE"
	r.Metadata = "init account"
	r.Type = 0

	res := ts.PrivateContractCall(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	// "hash":"278a1f189e235fee846baf22b4ff699e702f9dd407d2361bbfb41159d57f4a2f"
	fmt.Println("res: ", string(dataByte))
}

func TestEvaluateFee(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	var r request.BIFTransactionEvaluateFeeRequest
	senderAddresss := "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
	destAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
	bifAmount := 10

	operation := request.BIFGasSendOperation{
		DestAddress: destAddress,
		Amount:      int64(bifAmount),
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.GAS_SEND,
		},
	}
	r.SourceAddress = senderAddresss
	r.Operation = operation
	r.SignatureNumber = 1

	res := ts.EvaluateFee(r)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}

func TestBIFSubmit(t *testing.T) {
	ts := GetTransactionInstance(SDK_INSTANCE_URL)
	publicKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
	senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"

	destAddress := "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
	amount := 10
	operation := request.BIFGasSendOperation{
		DestAddress: destAddress,
		Amount:      int64(amount),
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.GAS_SEND,
		},
	}

	// 一、获取交易发起的账号nonce值
	accountGetInfoUrl := common.AccountGetInfoURL(ts.url, senderAddress)
	dataByte, err := http.HttpGet(accountGetInfoUrl)
	if err != nil {
		t.Error(err)
	}

	var nonceResponse response.BIFAccountGetNonceResponse
	err = json.Unmarshal(dataByte, &nonceResponse)
	if err != nil {
		t.Error(err)
	}
	if nonceResponse.ErrorCode != common.SUCCESS {
		t.Error(err)
	}
	nonce := nonceResponse.Result.Nonce

	// 二、构建操作、序列化交易
	// 初始化请求参数 BIFTransactionSerializeRequest
	serializeRequest := request.BIFTransactionSerializeRequest{
		SourceAddress: senderAddress,
		Nonce:         nonce + 1,
		GasPrice:      common.GAS_PRICE,
		FeeLimit:      common.FEE_LIMIT,
		Operation:     operation,
	}
	// BIFTransactionSerializeResponse
	serializeResponse := ts.BIFSerializable(serializeRequest)
	if serializeResponse.ErrorCode != common.SUCCESS {
		t.Error(err)
	}
	transactionBlob := serializeResponse.Result.TransactionBlob
	blob, err := hex.DecodeString(transactionBlob)
	if err != nil {
		t.Error(err)
	}
	// 三、签名
	signData, err := key.Sign([]byte(publicKey), []byte(blob))
	if err != nil {
		t.Error(err)
	}

	pubKey, err := key.GetEncPublicKey([]byte(publicKey))
	if err != nil {
		t.Error(err)
	}

	submitRequest := request.BIFTransactionSubmitRequest{
		Serialization: hex.EncodeToString(blob),
		SignData:      hex.EncodeToString(signData),
		PublicKey:     pubKey,
	}

	res := ts.BIFSubmit(submitRequest)
	if res.ErrorCode != 0 {
		t.Error(res.ErrorDesc)
	}

	dataByte, err = json.Marshal(res)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", string(dataByte))
}
