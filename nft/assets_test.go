package nft

import (
	"context"
	"github.com/rovergulf/eth-contracts-go/tests"
)

func (suite *NFTTestSuite) TestFindCollectionOwnerAssets() {
	ctx := context.Background()

	assets, err := FindCollectionOwnerAssets(ctx, suite.backend, AssetsRequest{
		Address: suite.erc721Address,
		Owner:   tests.Account0,
		//LoadMetadata: false,
	})
	suite.Nil(err, "Should not return an error")
	suite.NotNil(assets, "Should return valid results finding owner assets")
}
