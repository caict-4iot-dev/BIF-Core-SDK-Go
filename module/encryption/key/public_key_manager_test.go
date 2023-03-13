package key

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetPublicKeyManagerEd25519(t *testing.T) {
	//privateKey := "priSPKt2hVqYmUoyJ24vYb8uzhu7VCab9X9BHz8K68sQZ61Yy5"
	//publicKey := "b06566da9b97a860598e14469cb1dd3af1bf6c5cb88d549379d49d83b5b0bb23d63bff"
	//address:= "did:bid:efdPP4ympm36RxTYQJPber1PRKmLWZE"
	//encPrivateKey := "priSPKkrpK2RoQwtVqN1byRRFJjspxYnjkrYei8439SSytpLj6"
	encPrivateKey := "priSPKt2hVqYmUoyJ24vYb8uzhu7VCab9X9BHz8K68sQZ61Yy5"
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), ED25519)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("EncPublicKey: ", publicKeyManager.EncPublicKey)
	fmt.Println("EncAddress: ", publicKeyManager.EncAddress)
}
func TestGetPublicKeyManagerByPublicKey(t *testing.T) {

	encPublicKey := "b06566ef131e8b1d223f5c3e89558de82b888c1cd5fa0d6c940458e9f6309040cfb28f"
	publicKeyManager, err := GetPublicKeyManagerByPublicKey(encPublicKey)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("EncPublicKey: ", publicKeyManager.EncPublicKey)
	fmt.Println("EncAddress: ", publicKeyManager.EncAddress)
	fmt.Println("rawPublicKey", publicKeyManager.RawPublicKey)
}
func TestGetPublicKeyManagerSm2(t *testing.T) {

	// did:bid:zfAKvFU7N6SazT3DsDG8YMZ13oxFF9ba
	//encPrivateKey := "priSrrUcLq9WQC1vVCi5deuASW1jph6bXEieQ1f2nwCzJXWTte"
	encPrivateKey := "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), SM2)
	if err != nil {
		t.Error(err)
	}
	keyType, rawPrivateKey, err := GetRawPrivateKey([]byte(encPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("EncPublicKey: ", publicKeyManager.EncPublicKey)
	fmt.Println("rawPublicKey: ", hex.EncodeToString(publicKeyManager.RawPublicKey))
	fmt.Println("EncAddress: ", publicKeyManager.EncAddress)
	fmt.Println("rawPrivateKey: ", hex.EncodeToString(rawPrivateKey))
	fmt.Println("keyType: ", keyType)
	// 5e221667893cb1cbdf31c60f76b7c304446d8114e9b6173ee45aa40bd072acd9
}

func TestSignEd25519(t *testing.T) {
	encPrivateKey := "priSPKt8AFKhaWpcYKZxGFJfpZJahKTpDrUwDV4k7myHkLzRje"
	msg := "hello word"
	// 签名
	signMsg, err := Sign([]byte(encPrivateKey), []byte(msg))
	if err != nil {
		t.Error(err)
	}
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), ED25519)
	if err != nil {
		t.Error(err)
	}

	// 验签
	isOK := Verify([]byte(publicKeyManager.EncPublicKey), []byte(msg), signMsg)
	if !isOK {
		t.Error("verify sign message is failed")
	}

	fmt.Println("result:", isOK)
}

func TestSignSm2(t *testing.T) {
	//encPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
	encPrivateKey := "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
	msg := "hello word"
	// 签名
	signMsg, err := Sign([]byte(encPrivateKey), []byte(msg))
	if err != nil {
		t.Error(err)
	}

	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), SM2)
	if err != nil {
		t.Error(err)
	}

	// 验签
	isOK := Verify([]byte(publicKeyManager.EncPublicKey), []byte(msg), signMsg)
	if !isOK {
		t.Error("verify sign message is failed")
	}

	fmt.Println("result:", isOK)
}

func TestVerify_other(t *testing.T) {
	encPrivateKey := "priSPKt8AFKhaWpcYKZxGFJfpZJahKTpDrUwDV4k7myHkLzRje"
	msg := "0a286469643a6269643a65667a51313974516d5a56384256365045374868756f563866386264754e76591001223f0807523b0a286469643a6269643a65664e69515045476e68545071614661746f463170397767723135325036384610011a0d7b22626172223a22666f6f227d2a080123456789abcdef30c0843d3801"
	signMsg := "59e2d313be2016a35724c6653e488cd938551726a9b475ed44f9650001db7ba65c6acc4ca375b841f6bb02a1f29bb4085083ff516a3d336872840b655f857a05"
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), ED25519)
	fmt.Println(err)
	signMsgByte, err := hex.DecodeString(signMsg)
	msgByte, err := hex.DecodeString(msg)
	isOK := Verify([]byte(publicKeyManager.EncPublicKey), msgByte, signMsgByte)
	fmt.Println(isOK)
}

func TestVerify_mismatching(t *testing.T) {
	encPrivateKey := "priSPKt8AFKhaWpcYKZxGFJfpZJahKTpDrUwDV4k7myHkLzRje"
	msg := "0a286469643a6269643a65667a51313974516d5a56384256365045374868756f563866386264754e76591001223f0807523b0a286469643a6269643a65664e69515045476e68545071614661746f463170397767723135325036384610011a0d7b22626172223a22666f6f227d2a080123456789abcdef30c0843d3801"
	//signMsg := "59e2d313be2016a35724c6653e488cd938551726a9b475ed44f9650001db7ba65c6acc4ca375b841f6bb02a1f29bb4085083ff516a3d336872840b655f857a05"
	signMsg, err := Sign([]byte(encPrivateKey), []byte(msg))
	keyPair, err := GetBidAndKeyPairBySM2()
	fmt.Println(err)
	publicKeyErr := keyPair.GetEncPublicKey()
	//publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), ED25519)

	isOK := Verify([]byte(publicKeyErr), []byte(msg), signMsg)
	fmt.Println(isOK)
}
