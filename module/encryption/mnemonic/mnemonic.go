package mnemonic

import (
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"log"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip32"

	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonicCode ...
func GenerateMnemonicCode() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

func GeneratePrivateKeys(mnemonic string, hdPaths string, keyType int) (string, error) {

	hdPath := strings.Replace(hdPaths, "'", "", -1)
	path := strings.Split(hdPath, "/")
	seed := bip39.NewSeed(mnemonic, "")
	rootPrivateKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", err
	}

	path01, err := strconv.Atoi(path[1])
	if err != nil {
		return "", err
	}
	path02, err := strconv.Atoi(path[2])
	if err != nil {
		return "", err
	}
	path03, err := strconv.Atoi(path[3])
	if err != nil {
		return "", err
	}
	path04, err := strconv.Atoi(path[4])
	if err != nil {
		return "", err
	}
	path05, err := strconv.Atoi(path[5])
	if err != nil {
		return "", err
	}

	// Get BIP44 Extended Private Key
	m44H, err := rootPrivateKey.NewChildKey(bip32.FirstHardenedChild + uint32(path01))
	if err != nil {
		return "", err
	}
	m44H526H, err := m44H.NewChildKey(bip32.FirstHardenedChild + uint32(path02))
	if err != nil {
		return "", err
	}
	m44H526H1H, err := m44H526H.NewChildKey(bip32.FirstHardenedChild + uint32(path03))
	if err != nil {
		return "", err
	}
	m44H526H1H0PrivateKey, err := m44H526H1H.NewChildKey(uint32(path04))
	if err != nil {
		return "", err
	}
	m44H526H1H0H0PrivateKey, err := m44H526H1H0PrivateKey.NewChildKey(uint32(path05))
	if err != nil {
		return "", err
	}
	privateKeyManager, err := key.GetPrivateKeyManagerByRand(m44H526H1H0H0PrivateKey.Key, keyType)
	if err != nil {
		return "", err
	}

	return privateKeyManager.EncPrivateKey, nil
}
