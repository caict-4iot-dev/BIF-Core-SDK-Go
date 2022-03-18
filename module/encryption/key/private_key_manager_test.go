package key

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetPrivateKeyManagerEd25519(t *testing.T) {
	privateKeyManager, err := GetPrivateKeyManager(ED25519)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("EncPrivateKey: ", privateKeyManager.EncPrivateKey)
	fmt.Println("RawPrivateKey: ", hex.EncodeToString(privateKeyManager.RawPrivateKey))
}

func TestGetPrivateKeyManagerSM2(t *testing.T) {
	privateKeyManager, err := GetPrivateKeyManager(SM2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("EncPrivateKey: ", privateKeyManager.EncPrivateKey)
	fmt.Println("RawPrivateKey: ", hex.EncodeToString(privateKeyManager.RawPrivateKey))
	fmt.Println("RawPublicKey: ", hex.EncodeToString(privateKeyManager.RawPublicKey))
}

func TestGetRawPrivateKey(t *testing.T) {
	encPrivate := "priSrrfnLfgfaQhTrpZyrMVpYRYTwYMRZkQwZQmvGZpoeQwRAB"
	_, rawPrivateKey, err := GetRawPrivateKey([]byte(encPrivate))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(rawPrivateKey))
}
