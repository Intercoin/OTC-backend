package app

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"otc-backend/config"
)

type SwapStore interface {
}

type App struct {
	logger *zap.Logger
	cfg    config.Config

	ethClient *ethclient.Client
	bscClient *ethclient.Client

	swapStore SwapStore

	isSync bool
}

func NewApp(logger *zap.Logger, cfg config.Config, swapStore SwapStore) (*App, error) {
	ethClient, err := ethclient.Dial(cfg.Eth.RpcURL)
	if err != nil {
		return nil, err
	}

	bscClient, err := ethclient.Dial(cfg.Bsc.RpcURL)
	if err != nil {
		return nil, err
	}

	return &App{
		logger:    logger,
		cfg:       cfg,
		swapStore: swapStore,
		ethClient: ethClient,
		bscClient: bscClient,
	}, nil
}
