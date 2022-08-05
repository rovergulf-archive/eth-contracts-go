package ethutils

import (
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

var (
	testStrEth      = "12"
	testStrWei      = "1230000000000000000"
	testStrFloatEth = "12.3"
)

type TestDenominationSuite struct {
	suite.Suite
}

func TestDenomination(t *testing.T) {
	suite.Run(t, new(TestDenominationSuite))
}

func (suite *TestDenominationSuite) TestEthToWei() {
	ebn := new(big.Int)
	ebn.SetString(testStrEth, 10)
	eth := EtherToWei(ebn)
	ethResult, _ := new(big.Int).SetString("12000000000000000000", 10)
	suite.Equal(eth, ethResult, "should be equal values")
}

func (suite *TestDenominationSuite) TestWeiToEth() {
	wbn := new(big.Int)
	wbn.SetString(testStrWei, 10)
	ethAmount := WeiToEther(wbn)
	suite.Equal(ethAmount, big.NewInt(1), "should return integer value")
}

func (suite *TestDenominationSuite) TestWeiToFloatEth() {
	floatEth, err := ParseBigFloat(testStrFloatEth)
	suite.Nil(err, "should not return error on parsing valid float value")
	suite.Equal(floatEth.String(), big.NewFloat(12.3).String(), "should return integer value")
}

func (suite *TestDenominationSuite) TestFloatEtherToWei() {
	floatEth, err := ParseBigFloat(testStrFloatEth)
	suite.Nil(err, "should not return error on parsing valid float value")
	suite.Equal(floatEth.String(), big.NewFloat(12.3).String(), "should return integer value")

	floatWei := FloatEtherToWei(floatEth)
	expectedWei, _ := new(big.Int).SetString("12300000000000000000", 10)
	suite.Equal(floatWei.String(), expectedWei.String(), "should be equal values")
}
