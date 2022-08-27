package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/eth-contracts-go/abis/introspection/erc165"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc1155"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc721"
	"github.com/rovergulf/eth-contracts-go/nft/model"
	"github.com/rovergulf/eth-contracts-go/pkg/interfaces"
)

//type NFTContract interface {
//	*erc721.Erc721 | *erc1155.Erc1155
//}

type AssetsRequest struct {
	Address      common.Address `json:"address"`
	Owner        common.Address `json:"owner"`
	LoadMetadata bool           `json:"load_metadata"`
}

func FindCollectionOwnerAssets(
	ctx context.Context,
	backend bind.ContractBackend,
	req AssetsRequest,
) (*model.NFTOwnershipResponse, error) {

	erc165Instance, err := erc165.NewErc165(req.Address, backend)
	if err != nil {
		return nil, err
	}

	isERC721, _ := erc165Instance.SupportsInterface(nil, interfaces.ERC721)
	isERC1155, _ := erc165Instance.SupportsInterface(nil, interfaces.ERC1155)

	if !isERC721 && !isERC1155 {
		return nil, ErrInvalidContractInterface
	}
	result := model.NFTOwnershipResponse{
		NFTContractBaseMetadata: model.NFTContractBaseMetadata{
			Address:   req.Address,
			Owner:     common.Address{},
			IsERC721:  isERC721,
			IsERC1155: isERC1155,
		},
		AssetsCount:         0,
		CollectionsCount:    0,
		AssetsPerCollection: nil,
	}
	//collections := make(map[common.Address]bool)

	if isERC721 {
		erc721Instance, err := erc721.NewErc721(req.Address, backend)
		if err != nil {
			return nil, err
		}

		logs, err := erc721Instance.FilterTransfer(nil, nil, []common.Address{req.Owner}, nil)
		if err != nil {
			return nil, err
		}

		for logs.Next() {
			result.AssetsCount++

		}
	} else {
		erc1155Instance, err := erc1155.NewErc1155(req.Address, backend)
		if err != nil {
			return nil, err
		}

		// actually, I am not sure if this is the right way, as it may contain erc20/777 interfaces
		// looks like other interfaces should not conflict
		singleTransfers, err := erc1155Instance.FilterTransferSingle(nil, nil, []common.Address{req.Owner}, nil)
		if err != nil {
			return nil, err
		}

		for singleTransfers.Next() {
			result.AssetsCount++
		}

		batchTransfers, err := erc1155Instance.FilterTransferBatch(nil, nil, []common.Address{req.Owner}, nil)
		if err != nil {
			return nil, err
		}

		for batchTransfers.Next() {
			result.AssetsCount++
		}
	}

	return &result, nil
}
