package key

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/base"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/hash"
	"log"
	"math/big"
	"strings"

	"github.com/ZZMarquis/gm/sm2"
)

type PublicKeyManager struct {
	EncAddress   string
	EncPublicKey string
	RawPublicKey []byte
}

func GetPublicKeyManager(encPrivateKey []byte, keyType int) (*PublicKeyManager, error) {

	encPublicKey, err := GetEncPublicKey(encPrivateKey)
	if err != nil {
		return nil, err
	}
	rawPublicKey := GetRawPublicKey([]byte(encPublicKey))
	encAddress := GetEncAddress(rawPublicKey, "", keyType)

	var pubManager PublicKeyManager
	pubManager.EncAddress = encAddress
	pubManager.EncPublicKey = encPublicKey
	pubManager.RawPublicKey = rawPublicKey

	return &pubManager, nil
}

// GetEncPublicKey 星火私钥获取星火公钥
func GetEncPublicKey(encPrivateKey []byte) (string, error) {
	keyType, rawPrivateKey, err := GetRawPrivateKey(encPrivateKey)
	if err != nil {
		return "", err
	}
	var rawPublicKey []byte

	switch keyType {
	case ED25519:
		rawPublicKey = rawPrivateKey[32:]
	case SM2:
		priKey, err := sm2.RawBytesToPrivateKey(rawPrivateKey)
		if err != nil {
			return "", err
		}
		pubKey := sm2.CalculatePubKey(priKey)
		rawPublicKey, err = hex.DecodeString("04" + hex.EncodeToString(pubKey.GetRawBytes()))
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("type does not exist")
	}

	return EncPublicKey(rawPublicKey, keyType), nil
}

// EncPublicKey 原生公钥转星火公钥
func EncPublicKey(rawPublicKey []byte, keyType int) string {

	buff := make([]byte, len(rawPublicKey)+3)
	buff[0] = 0xB0
	switch keyType {
	case ED25519:
		buff[1] = ED25519_VALUE
	case SM2:
		buff[1] = SM2_VALUE
	default:
		return ""
	}

	buff[2] = BASE_58_VALUE
	buff = append(buff[:3], rawPublicKey...)

	return hex.EncodeToString(buff)
}

// GetEncAddress 获取地址
func GetEncAddress(rawPublicKey []byte, chainCode string, keyType int) string {

	hashPkey := hash.GenerateHashHex(rawPublicKey, keyType)
	encAddress := base.Base58Encode(hashPkey[10:])
	if chainCode == "" {
		switch keyType {
		case ED25519:
			return "did:bid:" + "ef" + encAddress

		case SM2:
			return "did:bid:" + "zf" + encAddress

		default:
			return ""
		}
	} else {
		return "did:bid:" + chainCode + ":" + "ef" + encAddress
	}
}

// GetRawPublicKey 星火公钥获取原生公钥
func GetRawPublicKey(encPublicKey []byte) []byte {

	rawPublicKey, err := hex.DecodeString(string(encPublicKey))
	if err != nil {
		return nil
	}

	return rawPublicKey[3:]
}

// IsAddressValid check address
func IsAddressValid(encAddress string) bool {
	err := encAddressValid(encAddress)
	if err != nil {
		log.Println("IsAddressValid is failed, err: ", err)
		return false
	}

	return true
}

func encAddressValid(encAddress string) error {

	if encAddress == "" {
		return errors.New("invalid address")
	}
	items := strings.Split(encAddress, ":")
	if len(items) != 3 && len(items) != 4 {
		return errors.New("invalid address")
	}
	if len(items) == 3 {
		encAddress = items[2]
	} else {
		encAddress = items[3]
	}

	prefix := string([]byte(encAddress)[:2])
	if !(prefix == "ef") && !(prefix == "zf") {
		return errors.New("invalid address")
	}

	address := []byte(encAddress)[2:]
	// base58时先校验字符串是否合法
	for _, a := range address {
		if !strings.Contains("123456789AbCDEFGHJKLMNPQRSTuVWXYZaBcdefghijkmnopqrstUvwxyz", string(a)) {
			return errors.New("invalid address")
		}
	}
	base58Address := base.Base58Decode(address)
	if len(base58Address) != 22 {
		return errors.New("invalid address")
	}

	return nil
}

func Verify(encPublicKey []byte, msg []byte, signMsg []byte, keyType int) bool {

	var isOK bool
	rawPublicKey := GetRawPublicKey(encPublicKey)
	switch keyType {
	case ED25519:
		isOK = ed25519.Verify(rawPublicKey, msg, signMsg)
	case SM2:
		publicKeyHex := hex.EncodeToString(rawPublicKey)
		publicKey, err := hex.DecodeString(publicKeyHex[2:])
		if err != nil {
			return false
		}
		pubKey, _ := sm2.RawBytesToPublicKey(publicKey)
		r := new(big.Int).SetBytes(signMsg[:32])
		s := new(big.Int).SetBytes(signMsg[32:])
		isOK = sm2.VerifyByRS(pubKey, []byte("1234567812345678"), msg, r, s)
	}

	return isOK
}

// GetPublicKeyManagerByPublicKey
func GetPublicKeyManagerByPublicKey(encPublicKey string) (*PublicKeyManager, error) {
	pblicKeyHex, err := hex.DecodeString(encPublicKey)
	if err != nil {
		return nil, errors.New("public key (" + encPublicKey + ") is invalid, please check")
	}
	var keyType = ED25519
	// 判断算法类型
	if pblicKeyHex[1] == ED25519_VALUE {
		keyType = ED25519
	} else if pblicKeyHex[1] == SM2_VALUE {
		keyType = SM2
	} else {
		return nil, errors.New("public key (" + encPublicKey + ") is invalid, please check")
	}
	if pblicKeyHex[2] != BASE_58_VALUE {
		return nil, errors.New("public key (" + encPublicKey + ") is invalid, please check")
	}

	publicKey := GetRawPublicKey([]byte(encPublicKey))
	encAddress := GetEncAddress(publicKey, "", keyType)

	var keyManager PublicKeyManager
	keyManager.EncPublicKey = encPublicKey
	keyManager.EncAddress = encAddress

	return &keyManager, nil
}
