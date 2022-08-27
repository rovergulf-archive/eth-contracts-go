package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc20"
	"github.com/rovergulf/eth-contracts-go/pkg/ethutils"
	"github.com/rovergulf/eth-contracts-go/tests"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type DefiTestSuite struct {
	suite.Suite

	auth *keystore.Key

	eth     *ethclient.Client
	backend *backends.SimulatedBackend

	erc20Address common.Address
	erc20        *erc20.Erc20
}

func (suite *DefiTestSuite) SetupSuite() {
	if len(tests.TestProviderUrl) > 0 {
		client, err := ethclient.DialContext(context.Background(), tests.TestProviderUrl)
		if err != nil {
			log.Fatalf("Unable to connect test network: %s", err)
		}
		suite.eth = client
	}

	suite.backend = tests.NewFakeBackend()

	key, err := ethutils.PrivateKeyStringToKey(tests.PrivateKey0)
	if err != nil {
		log.Fatal(err)
	}
	suite.auth = key

	deployer, err := bind.NewKeyedTransactorWithChainID(suite.auth.PrivateKey, suite.backend.Blockchain().Config().ChainID)
	if err != nil {
		log.Fatal(err)
	}

	addr, _, instance, err := erc20.DeployErc20(deployer, suite.backend, "SomeToken", "ST")
	if err != nil {
		log.Fatalf("Unable to upload test erc20: %s", err)
	}
	suite.backend.Commit()

	suite.erc20Address = addr
	suite.erc20 = instance
}

func (suite *DefiTestSuite) TearDownSuite() {
	if suite.eth != nil {
		suite.eth.Close()
	}

	suite.backend.Close()
}

func TestDefiTestSuite(t *testing.T) {
	suite.Run(t, new(DefiTestSuite))
}
