package nft

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/eth-contracts-go/nft/model"
)

func FindCollectionOwnerAssets(
	ctx context.Context,
	backend bind.ContractBackend,
	address common.Address,
	owner common.Address,
) (*model.NFTOwnershipResponse, error) {
	// actually, I am not sure if this is the right way, as it may contain erc20/777 interfaces
	// looks like other interfaces should not conflict
	logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{address},
		//Topics:    [][]common.Hash{{model.TransferTopicId}},

	})
	if err != nil {
		return nil, err
	}

	result := new(model.NFTOwnershipResponse)
	//collections := make(map[common.Address]bool)

	for _, log := range logs {
		//
		fmt.Println(log.Address.String(), log.BlockNumber, log.TxHash.String())
		result.AssetsCount++
	}

	return result, nil
}
