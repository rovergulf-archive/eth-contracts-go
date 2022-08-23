package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc1155"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc721"
	"github.com/rovergulf/eth-contracts-go/pkg/ethutils"
	"github.com/rovergulf/eth-contracts-go/tests"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

var (
	fakeMetadataUri = "https://api.rovergulf.net/nft/metadata/test-chubbies/"
)

type NFTTestSuite struct {
	suite.Suite

	backend *backends.SimulatedBackend
	auth    *keystore.Key

	erc721        *erc721.Erc721
	erc721Address common.Address

	erc1155        *erc1155.Erc1155
	erc1155Address common.Address

	handler *Handler
}

func (suite *NFTTestSuite) SetupSuite() {

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

	erc721Addr, _, erc721Instance, err := erc721.DeployErc721(deployer, suite.backend, "SomeToken", "ST")
	if err != nil {
		log.Fatalf("Unable to upload test erc721: %s", err)
	}
	suite.backend.Commit()

	suite.erc721Address = erc721Addr
	suite.erc721 = erc721Instance

	erc1155Addr, _, erc1155Instance, err := erc1155.DeployErc1155(deployer, suite.backend, fakeMetadataUri)
	if err != nil {
		log.Fatalf("Unable to upload test erc1155: %s", err)
	}
	suite.backend.Commit()

	suite.erc1155Address = erc1155Addr
	suite.erc1155 = erc1155Instance

	if len(tests.TestProviderUrl) > 0 {
		nftHandler, err := NewHandler(context.Background(), tests.TestProviderUrl)
		if err != nil {
			log.Fatalf("Unable to init nft handler: %s", err)
		}
		suite.handler = nftHandler
	}
}

func (suite *NFTTestSuite) TearDownSuite() {
	suite.backend.Close()

	if suite.handler != nil {
		suite.handler.Shutdown(context.Background())
	}
}

func TestNFTTestSuite(t *testing.T) {
	suite.Run(t, new(NFTTestSuite))
}
