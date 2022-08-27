package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/eth-contracts-go/tests"
)

func (suite *DefiTestSuite) TestGetAddressBalance() {
	if suite.eth == nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	balance, err := GetAddressBalance(ctx, suite.eth, tests.Account0)
	suite.Nil(err, "should not return an error on balance check")
	suite.NotNil(balance, "balance result shouldn't be nil")
}

func (suite *DefiTestSuite) TestGetTokenAddressBalance() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	balance, err := GetTokenAddressBalance(ctx, suite.backend, suite.erc20Address, tests.Account0)
	suite.Nil(err, "should not return an error on token balance check")
	suite.NotNil(balance, "balance result shouldn't be nil")

	_, err2 := GetTokenAddressBalance(ctx, suite.backend, common.HexToAddress("0x0"), tests.Account0)
	suite.NotNil(err2, "should return error on non-existent contract or its bad address")
}
