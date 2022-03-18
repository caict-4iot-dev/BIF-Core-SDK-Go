package mnemonic

import (
	"fmt"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"testing"
)

func TestGenerateMnemonicCode(t *testing.T) {
	mnemonic, err := GenerateMnemonicCode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("mnemonic:", mnemonic)
}

func TestGeneratePrivateKeys(t *testing.T) {
	mnemonic := "style orchard science puppy place differ benefit thing wrap type build scare"
	hdPaths := "m/44'/526'/1'/0/0"
	keyType := key.ED25519
	encPrivateKey, err := GeneratePrivateKeys(mnemonic, hdPaths, keyType)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("encPrivateKey:", encPrivateKey)
}
