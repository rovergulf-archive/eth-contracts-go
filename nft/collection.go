package nft

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/eth-contracts-go/abis/introspection/erc165"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc721"
	"github.com/rovergulf/eth-contracts-go/nft/model"
	"github.com/rovergulf/eth-contracts-go/pkg/interfaces"
)

var ErrInvalidContractInterface = errors.New("invalid contract interface")

func FindCollectionMetadata(ctx context.Context, backend bind.ContractBackend, address common.Address) (*model.NFTCollectionResponse, error) {
	//lb, err := backend.BlockNumber(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//lastBlock := new(big.Int).SetUint64(lb)

	erc165Instance, err := erc165.NewErc165(address, backend)
	if err != nil {
		return nil, err
	}

	isERC721, err := erc165Instance.SupportsInterface(nil, interfaces.ERC721)
	if err != nil {
		return nil, err
	}
	isERC1155, err := erc165Instance.SupportsInterface(nil, interfaces.ERC1155)
	if err != nil {
		return nil, err
	}

	if !isERC721 && !isERC1155 {
		return nil, ErrInvalidContractInterface
	}

	result := &model.NFTCollectionResponse{
		Address:            address,
		CollectionMetadata: &model.CollectionMetadata{},
	}

	if isERC721 {
		erc721Instance, err := erc721.NewErc721(address, backend)
		if err != nil {
			return nil, err
		}

		name, err := erc721Instance.Name(nil)
		if err != nil {
			return nil, err
		}

		symbol, err := erc721Instance.Symbol(nil)
		if err != nil {
			return nil, err
		}

		result.InterfaceID = "ERC721"
		result.Name = name
		result.Symbol = symbol
	}

	return result, nil
}
