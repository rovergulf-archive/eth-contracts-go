package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc20"
	"math/big"
)

func GetAddressBalance(ctx context.Context, client *ethclient.Client, address common.Address) (*big.Int, error) {
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	bn := new(big.Int)
	bn.SetUint64(blockNumber)
	balance, err := client.BalanceAt(ctx, address, bn)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetTokenAddressBalance(ctx context.Context, client *ethclient.Client, token common.Address, address common.Address) (*big.Int, error) {
	erc20Instance, err := erc20.NewErc20(token, client)
	if err != nil {
		return nil, err
	}

	balance, err := erc20Instance.BalanceOf(nil, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
