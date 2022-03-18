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
