package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type NFTContractBaseMetadata struct {
	Address   common.Address `json:"address"`
	Owner     common.Address `json:"owner"`
	IsERC721  bool           `json:"is_erc_721"`
	IsERC1155 bool           `json:"is_erc_1155"`
}

type NFTOwnershipResponse struct {
	NFTContractBaseMetadata
	AssetsCount         int64                    `json:"assets_count"`
	CollectionsCount    int64                    `json:"collections_count"`
	AssetsPerCollection map[common.Address]int64 `json:"assets_per_collection"`
}

type NFTCollectionResponse struct {
	NFTContractBaseMetadata
	*CollectionMetadata
	InterfaceID string `json:"interface_id"`
}

func (c NFTCollectionResponse) String() string {
	return fmt.Sprintf("InterfaceID: %s\n"+
		"Name: %s\n"+
		"Symbol: %s\n"+
		"Owner: %s",
		c.InterfaceID, c.Name, c.Symbol, c.Owner)
}
