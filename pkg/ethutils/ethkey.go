package ethutils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

func init() {
	gob.Register(elliptic.P256())
}

func PrivateKeyStringToKey(pkString string) (*keystore.Key, error) {
	privateKey, err := crypto.HexToECDSA(pkString)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	return &keystore.Key{
		Id:         uuid.New(),
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA),
		PrivateKey: privateKey,
	}, nil
}

func NewEthRandomKey() (*keystore.Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	key := &keystore.Key{
		Id:         uuid.New(),
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}

	return key, nil
}

func SaveEthKeyToFile(key *keystore.Key, path string, auth string) error {
	payload, err := keystore.EncryptKey(key, auth, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	return os.WriteFile(path, payload, 0755)
}

func LoadEthKeyFromFile(path string, auth string) (*keystore.Key, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return keystore.DecryptKey(body, auth)
}
