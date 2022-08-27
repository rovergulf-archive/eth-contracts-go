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
	erc165Instance, err := erc165.NewErc165(address, backend)
	if err != nil {
		return nil, err
	}

	isERC721, _ := erc165Instance.SupportsInterface(nil, interfaces.ERC721)
	isERC1155, _ := erc165Instance.SupportsInterface(nil, interfaces.ERC1155)

	if !isERC721 && !isERC1155 {
		return nil, ErrInvalidContractInterface
	}

	result := &model.NFTCollectionResponse{
		NFTContractBaseMetadata: model.NFTContractBaseMetadata{
			Address:   address,
			Owner:     common.Address{},
			IsERC721:  isERC721,
			IsERC1155: isERC1155,
		},
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

		result.Name = name
		result.Symbol = symbol
	} else {
		//erc1155Instance, err := erc1155.NewErc1155(address, backend)
		//if err != nil {
		//	return nil, err
		//}

		//
	}

	return result, nil
}
