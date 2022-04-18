package app

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"otc-backend/types"
	"otc-backend/web3"
)

func (app *App) WithRetry(fn func() error, warn string) error {
	return retry.Do(fn,
		retry.Attempts(app.cfg.RetryAttempts),
		retry.Delay(app.cfg.RetryDelay*time.Millisecond),
		retry.OnRetry(func(n uint, err error) {
			app.logger.Warn(warn, zap.Uint("attempt", n), zap.Error(err))
		}))
}

func (app *App) GetTrade(tradehash string) (*types.Trade, error) {

	//TODO: implement me

	return nil, nil
}

func (app *App) Start(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(3 * time.Second):
				var err error
				app.ethClient, err = ethclient.Dial(app.cfg.Eth.RpcURL)
				if err != nil {
					app.logger.Error("failed to update client", zap.Error(err))
					continue
				}

				app.swap, err = web3.NewOTCSwap(common.HexToAddress(app.cfg.Eth.SwapAddress), app.ethClient)
				if err != nil {
					app.logger.Error("failed to generate nft abi", zap.Error(err))
					continue
				}

				lastBlockNumber, err := app.blockStore.GetLastBlock(ctx)
				if err != nil {
					app.logger.Error("failed to get last block number", zap.Error(err))
					continue
				}
				blockNumber, err := app.ethClient.BlockNumber(ctx)
				if err != nil {
					app.logger.Error("failed to get block number", zap.Error(err))
					continue
				}
				if blockNumber <= lastBlockNumber {
					continue
				}

				if blockNumber-lastBlockNumber > 1000 {
					blockNumber = lastBlockNumber + 1000
				}

				start := lastBlockNumber
				end := blockNumber
				if end-start > 1000 {
					end = start + 1000
				}
				for {
					app.logger.Info("getting new events...", zap.Uint64("start", start), zap.Uint64("end", end))
					opts := &bind.FilterOpts{
						Start:   start,
						End:     &end,
						Context: ctx,
					}
					transfers, err := app.swap.FilterNewTransfer(opts)
					if err != nil {
						app.logger.Error("failed to filter swaps transfers", zap.Error(err))
						continue
					}
					for transfers.Next() {
						tradehash := transfers.Event.Poster.String()
						err := app.SyncTrade(ctx, tradehash)
						if err != nil {
							app.logger.Error("failed to sync trade", zap.Error(err), zap.String("tradehash", tradehash))
							continue
						}
					}
					_ = transfers.Close()

					if end >= blockNumber {
						break
					}
					start = end
					end = start + 1000
					if end > blockNumber {
						end = blockNumber
					}
				}

				_, err = app.blockStore.Update(ctx, blockNumber)
				if err != nil {
					app.logger.Error("failed to update last block number", zap.Error(err))
					continue
				}

				app.ethClient.Close()

			case <-ctx.Done():
				return
			}
		}
	}(ctx)
}
