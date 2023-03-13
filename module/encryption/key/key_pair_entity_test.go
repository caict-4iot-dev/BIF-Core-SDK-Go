package key

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetBidAndKeyPairEd25519(t *testing.T) {

	keyPair, err := GetBidAndKeyPair()
	if err != nil {
		t.Error(err)
	}
	encAddress := keyPair.GetEncAddress()
	encPublicKey := keyPair.GetEncPublicKey()
	encPrivateKey := keyPair.GetEncPrivateKey()
	rawPublicKey := keyPair.GetRawPublicKey()
	rawPrivateKey := keyPair.GetRawPrivateKey()

	fmt.Println("encAddress: ", encAddress)
	fmt.Println("encPublicKey: ", encPublicKey)
	fmt.Println("encPrivateKey: ", encPrivateKey)
	fmt.Println("rawPublicKey: ", hex.EncodeToString(rawPublicKey))
	fmt.Println("rawPrivateKey: ", hex.EncodeToString(rawPrivateKey))
}

func TestGetBidAndKeyPairSm2(t *testing.T) {

	keyPair, err := GetBidAndKeyPairBySM2()
	if err != nil {
		t.Error(err)
	}
	encAddress := keyPair.GetEncAddress()
	encPublicKey := keyPair.GetEncPublicKey()
	encPrivateKey := keyPair.GetEncPrivateKey()
	rawPublicKey := keyPair.GetRawPublicKey()
	rawPrivateKey := keyPair.GetRawPrivateKey()

	fmt.Println("encAddress: ", encAddress)
	fmt.Println("encPublicKey: ", encPublicKey)
	fmt.Println("encPrivateKey: ", encPrivateKey)
	fmt.Println("rawPublicKey: ", hex.EncodeToString(rawPublicKey))
	fmt.Println("rawPrivateKey: ", hex.EncodeToString(rawPrivateKey))
}

func TestName(t *testing.T) {
	//encPublicKey:  b07a660402e887472af193e87fb2a22334440a6f9efbc8aa630a21de4b229ae1c3b5bfe0a9e353ad36fed07a0dea2453d0a6c91fb1130fea24aaec90fad98b02656549ec
	//rawPublicKey:  02e887472af193e87fb2a22334440a6f9efbc8aa630a21de4b229ae1c3b5bfe0a9e353ad36fed07a0dea2453d0a6c91fb1130fea24aaec90fad98b02656549ec
	rawPublicKey, err := hex.DecodeString("02e887472af193e87fb2a22334440a6f9efbc8aa630a21de4b229ae1c3b5bfe0a9e353ad36fed07a0dea2453d0a6c91fb1130fea24aaec90fad98b02656549ec")
	if err != nil {
		t.Error(err)
	}
	encPublicKey := EncPublicKey(rawPublicKey, SM2)
	fmt.Println(encPublicKey)
}
