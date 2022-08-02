package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc721"
	"go.uber.org/zap"
)

type ERC721Handler struct {
	client   *ethclient.Client
	logger   *zap.SugaredLogger
	contract *erc721.Erc721
}

func NewERC721Handler(
	client *ethclient.Client,
	address common.Address,
) (*ERC721Handler, error) {
	contract, err := erc721.NewErc721(address, client)
	if err != nil {
		return nil, err
	}

	return &ERC721Handler{
		client:   client,
		logger:   zap.S(),
		contract: contract,
	}, nil
}

func (h *ERC721Handler) WithLogger(logger *zap.SugaredLogger) *ERC721Handler {
	handler := h
	handler.logger = logger
	return handler
}
