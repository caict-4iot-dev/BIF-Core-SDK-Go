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
	encPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
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
	isOK := Verify([]byte(publicKeyManager.EncPublicKey), []byte(msg), signMsg, ED25519)
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
	isOK := Verify([]byte(publicKeyManager.EncPublicKey), []byte(msg), signMsg, SM2)
	if !isOK {
		t.Error("verify sign message is failed")
	}

	fmt.Println("result:", isOK)
}
