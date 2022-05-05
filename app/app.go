package app

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"otc-backend/repository/postgres"

	"otc-backend/config"
	"otc-backend/types"
	"otc-backend/web3"
)

type SwapStore interface {
	GetTrade(ctx context.Context, tradehash string) (*postgres.Trade, error)
	CreateTrade(ctx context.Context, trade *postgres.Trade) error
	UpdateTrade(ctx context.Context, tradeID int, network postgres.Network) error
	SaveLock(ctx context.Context, lock *postgres.Lock) error
	SaveEngage(ctx context.Context, lock *postgres.Engage) error
	SaveClaim(ctx context.Context, lock *postgres.Claim) error
}

type App struct {
	logger *zap.Logger
	cfg    config.Config

	ethClient *ethclient.Client
	bscClient *ethclient.Client

	swapEth *web3.OTCSwap
	swapBSC *web3.OTCSwap

	swapStore  SwapStore
	blockStore BlockStore

	isSync bool
}

func NewApp(logger *zap.Logger, cfg config.Config, swapStore SwapStore, blockStore BlockStore) (*App, error) {
	ethClient, err := ethclient.Dial(cfg.Eth.RpcURL)
	if err != nil {
		return nil, err
	}

	bscClient, err := ethclient.Dial(cfg.Bsc.RpcURL)
	if err != nil {
		return nil, err
	}

	return &App{
		logger:     logger,
		cfg:        cfg,
		swapStore:  swapStore,
		blockStore: blockStore,
		ethClient:  ethClient,
		bscClient:  bscClient,
	}, nil
}

type BlockStore interface {
	GetLastBlock(context.Context, postgres.Network) (uint64, error)
	Update(context.Context, postgres.Network, uint64) (uint64, error)
}

func (app *App) GetTrade(ctx context.Context, tradehash string) (*types.Trade, error) {
	trade, err := app.swapStore.GetTrade(ctx, tradehash)
	if err != nil {
		return nil, err
	}
	return types.TradeFromDB(trade), nil
}

func (app *App) CreateTrade(ctx context.Context, trade *types.Trade) error {
	return app.swapStore.CreateTrade(ctx, trade.ToDB())
}

func (app *App) UpdateTrade(ctx context.Context, tradeID int, network2 types.Network) error {
	return app.swapStore.UpdateTrade(ctx, tradeID, network2)
}

func (app *App) SetLock(ctx context.Context, lock *types.Lock) error {
	return app.swapStore.SaveLock(ctx, lock.ToDB())
}
func (app *App) SetEngage(ctx context.Context, engage *types.Engage) error {
	return app.swapStore.SaveEngage(ctx, engage.ToDB())
}
func (app *App) SetClaim(ctx context.Context, claim *types.Claim) error {
	return app.swapStore.SaveClaim(ctx, claim.ToDB())
}
