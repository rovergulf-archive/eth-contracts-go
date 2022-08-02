package api

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/abis/common/erc165"
	"go.uber.org/zap"
)

var nilInterfaceId = [4]byte{}

type ERC165Handler struct {
	client   *ethclient.Client
	logger   *zap.SugaredLogger
	contract *erc165.Erc165
}

func NewERC165Handler(
	client *ethclient.Client,
	address common.Address,
) (*ERC165Handler, error) {
	contract, err := erc165.NewErc165(address, client)
	if err != nil {
		return nil, err
	}

	return &ERC165Handler{
		client:   client,
		logger:   zap.S(),
		contract: contract,
	}, nil
}

func (h *ERC165Handler) WithLogger(logger *zap.SugaredLogger) *ERC165Handler {
	handler := h
	handler.logger = logger
	return handler
}

func (h *ERC165Handler) SupportsInterface(ctx context.Context, interfaceId [4]byte) (bool, error) {
	supports, err := h.contract.SupportsInterface(nil, interfaceId)
	if err != nil {
		h.logger.Errorw("ERC165: Unable to call supports interface method", "err", err)
		return false, err
	}

	return supports, nil
}
