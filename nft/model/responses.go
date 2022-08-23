package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type NFTOwnershipResponse struct {
	Owner               common.Address           `json:"owner"`
	AssetsCount         int64                    `json:"assets_count"`
	CollectionsCount    int64                    `json:"collections_count"`
	AssetsPerCollection map[common.Address]int64 `json:"assets_per_collection"`
}

type NFTCollectionResponse struct {
	InterfaceID string         `json:"interface_id"`
	Address     common.Address `json:"address"`
	Owner       common.Address `json:"owner"`
	*CollectionMetadata
}

func (c NFTCollectionResponse) String() string {
	return fmt.Sprintf("InterfaceID: %s\n"+
		"Name: %s\n"+
		"Symbol: %s\n"+
		"Owner: %s",
		c.InterfaceID, c.Name, c.Symbol, c.Owner)
}
