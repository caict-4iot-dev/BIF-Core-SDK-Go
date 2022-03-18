package key

type KeyPairEntity struct {
	EncAddress    string
	EncPublicKey  string
	EncPrivateKey string
	RawPrivateKey []byte
	RawPublicKey  []byte
}

func GetBidAndKeyPair() (*KeyPairEntity, error) {
	privateKeyManager, err := GetPrivateKeyManager(ED25519)
	if err != nil {
		return nil, err
	}
	publicKeyManager, err := GetPublicKeyManager([]byte(privateKeyManager.EncPrivateKey), ED25519)
	if err != nil {
		return nil, err
	}

	var keyPair KeyPairEntity
	keyPair.EncAddress = publicKeyManager.EncAddress
	keyPair.EncPublicKey = publicKeyManager.EncPublicKey
	keyPair.RawPublicKey = privateKeyManager.RawPublicKey
	keyPair.EncPrivateKey = privateKeyManager.EncPrivateKey
	keyPair.RawPrivateKey = privateKeyManager.RawPrivateKey

	return &keyPair, nil
}

func GetBidAndKeyPairBySM2() (*KeyPairEntity, error) {
	privateKeyManager, err := GetPrivateKeyManager(SM2)
	if err != nil {
		return nil, err
	}
	publicKeyManager, err := GetPublicKeyManager([]byte(privateKeyManager.EncPrivateKey), SM2)
	if err != nil {
		return nil, err
	}

	var keyPair KeyPairEntity
	keyPair.EncAddress = publicKeyManager.EncAddress
	keyPair.EncPublicKey = publicKeyManager.EncPublicKey
	keyPair.RawPublicKey = privateKeyManager.RawPublicKey
	keyPair.EncPrivateKey = privateKeyManager.EncPrivateKey
	keyPair.RawPrivateKey = privateKeyManager.RawPrivateKey

	return &keyPair, nil
}

func (k *KeyPairEntity) GetEncAddress() string {
	return k.EncAddress
}

func (k *KeyPairEntity) GetEncPublicKey() string {
	return k.EncPublicKey
}

func (k *KeyPairEntity) GetEncPrivateKey() string {
	return k.EncPrivateKey
}

func (k *KeyPairEntity) GetRawPublicKey() []byte {
	return k.RawPublicKey
}

func (k *KeyPairEntity) GetRawPrivateKey() []byte {
	return k.RawPrivateKey
}
