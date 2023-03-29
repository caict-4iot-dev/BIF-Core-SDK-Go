package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/base"
	"io"
	mrand "math/rand"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"
)

type KeyStore struct {
	Address    string
	AesctrIv   string
	CypherText string
	ScryptParams
	Version int
}

type ScryptParams struct {
	N    int
	P    int
	R    int
	Salt string
}

func GenerateKeyStore(encPrivateKey string, password string, n int, r int, p int, version int) (string, KeyStore) {

	var keyStore KeyStore
	dkLen := 32
	rand := mrand.Intn(10000000)
	randStr := strconv.Itoa(rand)
	key, err := scrypt.Key([]byte(password), []byte(randStr), n, r, p, dkLen)
	if err != nil {
		return "", keyStore
	}
	if encPrivateKey == "" {
		return "", keyStore
	}
	skeyTmp := base.Base58Decode([]byte(encPrivateKey))
	var keyType int
	switch skeyTmp[3] {
	case ED25519_VALUE:
		keyType = ED25519
		break
	case SM2_VALUE:
		keyType = SM2
		break
	}
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), keyType)
	if err != nil {
		return "", keyStore
	}
	keyStore.Address = publicKeyManager.EncAddress
	cypherText, aesctrIv := AESEncrypt(key, encPrivateKey)
	keyStore.CypherText = cypherText[len(aesctrIv):]
	keyStore.AesctrIv = aesctrIv
	var scryptParams ScryptParams
	scryptParams.N = n
	scryptParams.R = r
	scryptParams.P = p
	scryptParams.Salt = randStr
	keyStore.ScryptParams = scryptParams
	keyStore.Version = version

	return encPrivateKey, keyStore
}

func DecipherKeyStore(keyStore KeyStore, password string) string {
	dkLen := 32
	key, err := scrypt.Key([]byte(password), []byte(keyStore.Salt), keyStore.N, keyStore.R, keyStore.P, dkLen)
	if err != nil {
		return ""
	}
	var cypher string
	if strings.HasPrefix(keyStore.CypherText, keyStore.AesctrIv) {
		cypher = keyStore.CypherText
	} else {
		cypher = keyStore.AesctrIv + keyStore.CypherText
	}
	return AESDecrypt(key, cypher)
}

func AESEncrypt(key []byte, text string) (string, string) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ""
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), hex.EncodeToString(iv)
}

// AESDecrypt Decrypt from base64 to decrypted string
func AESDecrypt(key []byte, cryptoText string) string {
	ciphertext, _ := hex.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}
