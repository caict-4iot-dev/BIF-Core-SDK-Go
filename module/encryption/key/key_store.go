package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	mrand "math/rand"
	"strconv"

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
	rand := mrand.Intn(10000000)
	randStr := strconv.Itoa(rand)
	key, err := scrypt.Key([]byte(password), []byte(randStr), n, r, p, version)
	if err != nil {
		return "", keyStore
	}
	publicKeyManager, err := GetPublicKeyManager([]byte(encPrivateKey), ED25519)
	if err != nil {
		return "", keyStore
	}
	keyStore.Address = publicKeyManager.EncAddress
	cypherText, aesctrIv := AESEncrypt(key, encPrivateKey)
	keyStore.CypherText = cypherText
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
	key, err := scrypt.Key([]byte(password), []byte(keyStore.Salt), keyStore.N, keyStore.R, keyStore.P, keyStore.Version)
	if err != nil {
		return ""
	}
	return AESDecrypt(key, keyStore.CypherText)
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

	return fmt.Sprintf("%s", ciphertext)
}
