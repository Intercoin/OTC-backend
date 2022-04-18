package app

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"otc-backend/config"
	"otc-backend/types"
	"otc-backend/web3"
)

type SwapStore interface {
}

type App struct {
	logger *zap.Logger
	cfg    config.Config

	ethClient *ethclient.Client
	bscClient *ethclient.Client

	swap *web3.OTCSwap

	swapStore  SwapStore
	blockStore BlockStore

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

type BlockStore interface {
	GetLastBlock(context.Context) (uint64, error)
	Update(context.Context, uint64) (uint64, error)
}

func (app *App) UpdateTrade(ctx context.Context, trade *types.Trade) error {

	//TODO: implement me

	return nil
}
