package client

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var ErrEmptyProviderUrl = errors.New("empty url provider")

type EthClient struct {
	*ethclient.Client
	logger *zap.SugaredLogger
}

func New(providerUrl string) (*EthClient, error) {
	if len(providerUrl) == 0 {
		return nil, ErrEmptyProviderUrl
	}

	cl, err := ethclient.DialContext(context.Background(), providerUrl)
	if err != nil {
		return nil, err
	}

	return &EthClient{Client: cl, logger: zap.S()}, nil
}

func NewWithLogger(providerUrl string, logger *zap.SugaredLogger) (*EthClient, error) {
	if len(providerUrl) == 0 {
		return nil, ErrEmptyProviderUrl
	}

	cl, err := ethclient.DialContext(context.Background(), providerUrl)
	if err != nil {
		return nil, err
	}

	return &EthClient{
		Client: cl,
		logger: logger,
	}, nil
}
