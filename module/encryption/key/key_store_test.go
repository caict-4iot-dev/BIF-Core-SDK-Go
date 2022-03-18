package key

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateKeyStore(t *testing.T) {

	encPrivateKey := "priSPKrR4w6H89kRXaC2XZT5Lmj7XoCoBdvTuv7ySXSCDDGaZZ"
	password := "123456"
	n := 16384
	r := 8
	p := 1
	version := 32
	encPrivateKey, keyStore := GenerateKeyStore(encPrivateKey, password, n, r, p, version)
	dataByte, err := json.Marshal(keyStore)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("encPrivateKey: ", encPrivateKey)
	fmt.Println("keyStore: ", string(dataByte))
}

func TestDecipherKeyStore(t *testing.T) {
	// {"Address":"did:bid:ef24oGV9p46o1uwm2aFgZpScwW13r8nbu","AesctrIv":"978071003b6d24f0b861048cbd4c008b","CypherText":"978071003b6d24f0b861048cbd4c008bdcc773c2dfbf58ab5e9fd01fc0893b342e387d701058d9f7f9e769ee5f36da750c598213b46b1fa6157e9e80e55a33e75a83","N":16384,"P":1,"R":8,"Salt":"8498081","Version":32}
	keyStore := KeyStore{
		Address:    "did:bid:ef24oGV9p46o1uwm2aFgZpScwW13r8nbu",
		AesctrIv:   "978071003b6d24f0b861048cbd4c008b",
		CypherText: "978071003b6d24f0b861048cbd4c008bdcc773c2dfbf58ab5e9fd01fc0893b342e387d701058d9f7f9e769ee5f36da750c598213b46b1fa6157e9e80e55a33e75a83",
		ScryptParams: ScryptParams{
			N:    16384,
			P:    1,
			R:    8,
			Salt: "8498081",
		},
		Version: 32,
	}
	password := "123456"
	encPrivateKey := DecipherKeyStore(keyStore, password)
	fmt.Println("encPrivateKey:", encPrivateKey)
}
