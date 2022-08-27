package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

func (suite *NFTTestSuite) TestFindCollectionMetadata() {
	ctx := context.Background()

	info, err := FindCollectionMetadata(ctx, suite.backend, suite.erc721Address)
	suite.Nil(err, "Should be nil error")
	suite.NotNil(info, "Should successfully return collection metadata")
	if info != nil {
		suite.Equal(info.Name, testCollectionName, "Found collection name should match with test value")
	}

	_, errInvalidInterface := FindCollectionMetadata(ctx, suite.backend, common.HexToAddress("0x0"))
	suite.Equal(ErrInvalidContractInterface, errInvalidInterface, "should return invalid contract interface error")
}
