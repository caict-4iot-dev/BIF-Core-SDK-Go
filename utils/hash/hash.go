package hash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ZZMarquis/gm/sm3"
)

const (
	SHA256 = iota + 1
	SM3
)

func GenerateHashHex(src []byte, hashType int) []byte {
	var hashHex []byte
	switch hashType {
	case SHA256:
		hash := sha256.New()
		hash.Write(src)
		hashHex = hash.Sum(nil)
	case SM3:
		prefix, _ := hex.DecodeString("04")
		prefix = append(prefix, src...)
		hash := sm3.New()
		hash.Write(prefix)
		hashHex = hash.Sum(nil)
	}

	return hashHex
}
