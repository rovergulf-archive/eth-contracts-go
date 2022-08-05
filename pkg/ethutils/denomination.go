package ethutils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
)

// ParseBigFloat parse string value to big.Float
func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

// FloatEtherToWei formats Eth value argument to Wei big.Float
func FloatEtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func EtherToWei(eth *big.Int) *big.Int {
	return eth.Mul(eth, big.NewInt(params.Ether))
}

func WeiToEther(eth *big.Int) *big.Int {
	return eth.Div(eth, big.NewInt(params.Ether))
}

func WeiToFloatEther(eth *big.Int) *big.Float {
	ethAmount := new(big.Float)
	ethAmount.SetString(eth.String())
	ethAmount.Mul(ethAmount, big.NewFloat(params.Ether))
	return ethAmount
}
