package defi

import (
	"context"
	"github.com/rovergulf/eth-contracts-go/tests"
)

//func (suite *DefiTestSuite) TestGetAddressBalance() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	balance, err := GetAddressBalance(ctx, suite.backend, tests.Account0)
//}

func (suite *DefiTestSuite) TestGetTokenAddressBalance() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	balance, err := GetTokenAddressBalance(ctx, suite.backend, suite.erc20Address, tests.Account0)
	suite.Nil(err, "should not return an error on token balance check")
	suite.NotNil(balance, "balance result shouldn't be nil")
}
