package ethutils

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rovergulf/eth-contracts-go/tests"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
)

const (
	testKeystoreAuth = "suite suite suite suite suite suite suite suite suite suite suite suite"
	testEthKeyName   = "../../tmp/ethutils_test_key"
)

func init() {
	if _, err := os.Stat("../../tmp"); err != nil {
		if err := os.Mkdir("../../tmp", 0755); err != nil {
			log.Fatalf("Unable to create temp dir: %s", err)
		}
	}
}

type TestPrivateKeysSuite struct {
	suite.Suite
}

func (suite *TestPrivateKeysSuite) TestPrivateKeyStringToKey() {
	k, err := PrivateKeyStringToKey(tests.PrivateKey0)
	suite.Nil(err, "should be nil error on private key from string")
	suite.NotNil(k, "key shouldn't be nil")
	suite.Equal(k.Address.Hex(), tests.Account0.Hex(), "recovered address must match")

	k2, err2 := PrivateKeyStringToKey(hexutil.Encode(crypto.Keccak256([]byte("definitely not a pk"))))
	suite.NotNil(err2, "should return an err")
	suite.Nil(k2, "key2 cant be recovered")
}

func (suite *TestPrivateKeysSuite) TestFileKeystore() {
	k, err := NewEthRandomKey()
	suite.Nil(err, "should generate new key with no error")
	suite.NotNil(k, "should generate new key")

	suite.Nil(SaveEthKeyToFile(k, testEthKeyName, testKeystoreAuth), "should successfully save keystore file")

	loaded, err2 := LoadEthKeyFromFile(testEthKeyName, testKeystoreAuth)
	suite.Nil(err2, "should be nil error on load keystore file")
	suite.NotNil(loaded, "should load stored key")

	suite.Equal(k.Address.Hex(), loaded.Address.Hex(), "generated and loaded keys must match")

	suite.Nil(os.Remove(testEthKeyName))
}

func TestTestPrivateKeysSuite(t *testing.T) {
	suite.Run(t, new(TestPrivateKeysSuite))
}
