package key

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/ZZMarquis/gm/sm2"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/base"
)

const (
	// ED25519算法
	ED25519 = iota + 1
	// SM2算法
	SM2

	// 加密类型关键字
	ED25519_VALUE = 'e'
	SM2_VALUE     = 'z'
	// 字符编码类型关键字
	BASE_58_VALUE = 'f'
)

type PrivateKeyManager struct {
	EncPrivateKey string // 星火私钥
	RawPrivateKey []byte // 原生私钥
	RawPublicKey  []byte // 原生公钥
	TypeKey       string // 加密类型
}

func GetPrivateKeyManager(keyType int) (*PrivateKeyManager, error) {

	var rawPublicKey, rawPrivateKey []byte
	var encPrivateKey string
	var typeKey string
	switch keyType {
	case ED25519:
		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return nil, err
		}

		rawPrivateKey = privateKey[:32]
		rawPublicKey = publicKey
		typeKey = "ed25519"
	case SM2:
		privateKey, publicKey, err := sm2.GenerateKey(rand.Reader)
		if err != nil {
			return nil, err
		}
		rawPrivateKey = privateKey.GetRawBytes()
		rawPublicKey = publicKey.GetRawBytes()
		typeKey = "sm2"
	default:
		return nil, errors.New("type does not exist")
	}
	encPrivateKey = GetEncPrivateKey(rawPrivateKey, keyType)

	var priManager PrivateKeyManager
	priManager.RawPrivateKey = rawPrivateKey
	priManager.EncPrivateKey = encPrivateKey
	priManager.RawPublicKey = rawPublicKey
	priManager.TypeKey = typeKey

	return &priManager, nil
}

// GetPrivateKeyManagerByPrivateKey 根据星火私钥获取私钥对象
func GetPrivateKeyManagerByPrivateKey(encPrivateKey string) (*PrivateKeyManager, error) {
	keyType, rawPrivateKey, err := GetRawPrivateKey([]byte(encPrivateKey))
	if err != nil {
		return nil, err
	}
	var rawPublicKey []byte
	var typeKey string

	switch keyType {
	case ED25519:
		rawPublicKey = rawPrivateKey[32:]
		typeKey = "ed25519"
	case SM2:
		priKey, err := sm2.RawBytesToPrivateKey(rawPrivateKey)
		if err != nil {
			return nil, err
		}
		pubKey := sm2.CalculatePubKey(priKey)
		rawPublicKey, err = hex.DecodeString("04" + hex.EncodeToString(pubKey.GetRawBytes()))
		if err != nil {
			return nil, err
		}
		typeKey = "sm2"
	default:
		return nil, errors.New("unknown privateKey")
	}
	var priManager PrivateKeyManager
	priManager.RawPrivateKey = rawPrivateKey[:32]
	priManager.EncPrivateKey = encPrivateKey
	priManager.RawPublicKey = rawPublicKey
	priManager.TypeKey = typeKey

	return &priManager, nil
}

func GetPrivateKeyManagerByRand(seed []byte, keyType int) (*PrivateKeyManager, error) {

	var rawPublicKey, rawPrivateKey []byte
	var encPrivateKey string
	var typeKey string
	switch keyType {
	case ED25519:
		publicKey, privateKey, err := ed25519.GenerateKey(bytes.NewReader(seed))
		if err != nil {
			return nil, err
		}

		rawPrivateKey = privateKey[:32]
		rawPublicKey = publicKey
		typeKey = "ed25519"
	case SM2:
		privateKey, publicKey, err := sm2.GenerateKey(bytes.NewReader(seed))
		if err != nil {
			return nil, err
		}
		rawPrivateKey = privateKey.GetRawBytes()
		rawPublicKey = publicKey.GetRawBytes()
		typeKey = "sm2"
	default:
		return nil, errors.New("type does not exist")
	}
	encPrivateKey = GetEncPrivateKey(rawPrivateKey, keyType)

	var priManager PrivateKeyManager
	priManager.RawPrivateKey = rawPrivateKey
	priManager.EncPrivateKey = encPrivateKey
	priManager.RawPublicKey = rawPublicKey
	priManager.TypeKey = typeKey

	return &priManager, nil
}

// GetEncPrivateKey 原生私钥转星火私钥
func GetEncPrivateKey(privateKey []byte, keyType int) string {

	buff := make([]byte, len(privateKey)+5)
	buff[0] = 0x18
	buff[1] = 0x9E
	buff[2] = 0x99

	switch keyType {
	case ED25519:
		buff[3] = ED25519_VALUE
	case SM2:
		buff[3] = SM2_VALUE
	default:
		return ""
	}

	buff[4] = BASE_58_VALUE
	buff = append(buff[:5], privateKey...)

	return base.Base58Encode(buff)
}

// GetRawPrivateKey 星火私钥转原生私钥
func GetRawPrivateKey(encPrivateKey []byte) (int, []byte, error) {
	priKeyTmp := base.Base58Decode(encPrivateKey)
	if len(priKeyTmp) <= 5 {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 1")
	}

	if priKeyTmp[3] != ED25519_VALUE && priKeyTmp[3] != SM2_VALUE {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 2")
	}
	var keyType int
	switch priKeyTmp[3] {
	case ED25519_VALUE:
		{
			keyType = ED25519
		}
	case SM2_VALUE:
		{
			keyType = SM2
		}
	default:
		return 0, nil, errors.New("Private key (" + string(encPrivateKey) + ") is invalid")
	}
	if priKeyTmp[4] != BASE_58_VALUE {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 3")
	}

	var rawPrivateKey []byte
	switch keyType {
	case ED25519:
		rawPrivateKey = ed25519.NewKeyFromSeed(priKeyTmp[5:])
	case SM2:
		rawPrivateKey = priKeyTmp[5:]
	}

	return keyType, rawPrivateKey, nil
}

func Sign(encPrivate []byte, message []byte) ([]byte, error) {
	keyType, rawPrivateKey, err := GetRawPrivateKey(encPrivate)
	if err != nil {
		return nil, err
	}

	var sign []byte
	switch keyType {
	case ED25519:
		sign25519 := ed25519.Sign(rawPrivateKey, message)
		sign = sign25519
	case SM2:
		priKey, err := sm2.RawBytesToPrivateKey(rawPrivateKey)
		if err != nil {
			return nil, err
		}
		r, s, err := sm2.SignToRS(priKey, []byte("1234567812345678"), message)
		if err != nil {
			return nil, err
		}
		rBytes := r.Bytes()
		sBytes := s.Bytes()
		sig := make([]byte, 64)
		if len(rBytes) == 33 {
			copy(sig[:32], rBytes[1:])
		} else {
			copy(sig[:32], rBytes[:])
		}
		if len(sBytes) == 33 {
			copy(sig[32:], sBytes[1:])
		} else {
			copy(sig[32:], sBytes[:])
		}

		sign = sig
	}

	return sign, nil
}
