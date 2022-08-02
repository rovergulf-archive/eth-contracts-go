package ethutils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rovergulf/eth-contracts-go/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

type args struct {
	payload []byte
	sig     []byte
	address common.Address
}

type TestSigSuite struct {
	suite.Suite
	testArgs   []args
	key        *keystore.Key
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func (suite *TestSigSuite) SetupSuite() {
	key, err := PrivateKeyStringToKey(tests.PrivateKey0)
	if err != nil {
		suite.Fail(err.Error(), "should restore test key")
	}
	suite.key = key

	publicKeyBytes := crypto.FromECDSAPub(suite.key.PrivateKey.Public().(*ecdsa.PublicKey))
	addr := common.BytesToAddress(publicKeyBytes)
	fmt.Println(addr.Hex()) // 0x8A5C1b40b192451367f28E0088dd75E15de40C05

	data1 := []byte("hello")
	data2 := []byte("0x064f3961c7f78ac305e556ac8d22a36e82a0ca50c45ca8a237000f998b4457fe")
	data3 := []byte("hello world")
	hash1 := crypto.Keccak256Hash(data1)
	//fmt.Println(hash1.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
	hash2 := crypto.Keccak256Hash(data2)
	//fmt.Println(hash2.Hex()) // 0xf7cfaa8154ec5e25d0d05272e7ce38789d6566f232072df065911613e33ceccb
	hash3 := crypto.Keccak256Hash(data3)

	signature1, err := crypto.Sign(hash1.Bytes(), suite.key.PrivateKey)
	if err != nil {
		suite.Fail(err.Error())
	}

	signature2, err := crypto.Sign(hash2.Bytes(), suite.key.PrivateKey)
	if err != nil {
		suite.Fail(err.Error())
	}

	signature3, err := crypto.Sign(hash3.Bytes(), suite.key.PrivateKey)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.testArgs = append(suite.testArgs, args{
		payload: hash1.Bytes(),
		sig:     signature1,
		address: addr,
	}, args{
		payload: hash2.Bytes(),
		sig:     signature2,
		address: addr,
	}, args{
		payload: hash3.Bytes(),
		sig:     signature3,
		address: addr,
	}, args{
		payload: hash3.Bytes(),
		sig:     signature2,
		address: addr,
	})
}

func (suite *TestSigSuite) TearDownSuite() {}

func TestTestSigSuite(t *testing.T) {
	suite.Run(t, new(TestSigSuite))
}

func (suite *TestSigSuite) TestVerifySign() {
	tt1 := suite.testArgs[0]
	tt2 := suite.testArgs[1]
	tt3 := suite.testArgs[2]
	tt4 := suite.testArgs[3]
	suite.Nil(VerifySign(tt1.payload, tt1.sig, tt1.address), "should verify signature")
	suite.Nil(VerifySign(tt2.payload, tt2.sig, tt2.address), "should verify signature")
	suite.Nil(VerifySign(tt3.payload, tt3.sig, tt3.address), "should verify signature")
	suite.NotNil(VerifySign(tt4.payload, tt4.sig, tt4.address), "should not pass verification")
	suite.NotNil(VerifySign(tt3.payload, tt3.sig[:len(tt3.sig)-1], tt3.address), "should not pass verification")
}

func (suite *TestSigSuite) TestVerifyEcRecover() {
	tt1 := suite.testArgs[0]
	tt4 := suite.testArgs[3]
	suite.Nil(VerifyEcRecover(tt1.payload, tt1.sig, tt1.address), "should verify signature")
	suite.NotNil(VerifyEcRecover(tt4.payload, tt4.sig, tt4.address), "should not verify signature")
}

//func (suite *TestSigSuite) TestVerifySignNoRecoverId() {
//	tt1 := suite.testArgs[0]
//	tt4 := suite.testArgs[3]
//	suite.Nil(VerifySignNoRecoverId(crypto.Keccak256(tt1.payload), tt1.sig, tt1.address), "should verify signature")
//	suite.NotNil(VerifySignNoRecoverId(tt4.payload, tt4.sig, tt4.address), "should not verify signature")
//	suite.NotNil(VerifySignNoRecoverId(tt1.payload, tt1.sig[:len(tt1.sig)-1], tt1.address), "should not pass verification")
//}
