package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/pkg/logutils"
	"github.com/rovergulf/eth-contracts-go/pkg/traceutils"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Handler struct {
	client *ethclient.Client
	auth   *keystore.Key
	logger *zap.SugaredLogger
	tracer trace.Tracer
}

func (h *Handler) Shutdown(ctx context.Context) {
	h.client.Close()
}

func NewHandler(ctx context.Context, providerUrl string) (*Handler, error) {
	h := new(Handler)

	logger, err := logutils.NewLogger()
	if err != nil {
		return nil, err
	}
	h.logger = logger

	if viper.IsSet(traceutils.CollectorUrlConfigKey) {
		tp, err := traceutils.NewJaegerProvider(viper.GetString(traceutils.CollectorUrlConfigKey))
		if err != nil {
			return nil, err
		}
		h.tracer = tp.Tracer("defi")
	}

	client, err := ethclient.DialContext(ctx, providerUrl)
	if err != nil {
		return nil, err
	}
	h.client = client

	return h, nil
}
