package app

import (
	"context"
	"fmt"
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

func (app *App) Start(ctx context.Context) {
	go func(ctx context.Context) {
		app.logger.Info("Started Ethereum watcher...")
		for {
			select {
			case <-time.After(13 * time.Second):
				var err error
				app.ethClient, err = ethclient.Dial(app.cfg.Eth.RpcURL)
				if err != nil {
					app.logger.Error("failed to update client", zap.Error(err))
					continue
				}

				app.swapEth, err = web3.NewOTCSwap(common.HexToAddress(app.cfg.Eth.SwapAddress), app.ethClient)
				if err != nil {
					app.logger.Error("failed to generate nft abi", zap.Error(err))
					continue
				}

				lastBlockNumber, err := app.blockStore.GetLastBlock(ctx, types.Ethereum)
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
					app.logger.Info("getting new events...", zap.Uint64("start", start), zap.Uint64("end", end),
						zap.String("Network", fmt.Sprint(types.Ethereum)))
					opts := &bind.FilterOpts{
						Start:   start,
						End:     &end,
						Context: ctx,
					}
					transfers, err := app.swapEth.FilterNewTransfer(opts)
					if err != nil {
						app.logger.Error("failed to filter swaps transfers", zap.Error(err))
						continue
					}
					for transfers.Next() {
						err := app.SyncLock(ctx, transfers.Event, types.Ethereum)
						if err != nil {
							app.logger.Error("failed to sync trade", zap.Error(err), zap.String("tradehash",
								transfers.Event.Poster.String()))
							continue
						}
						app.logger.Info("event catched", zap.String("Network", fmt.Sprint(types.Ethereum)),
							zap.String("Tradehash", common.BytesToHash(transfers.Event.TradeHash[:]).String()))
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

				_, err = app.blockStore.Update(ctx, types.Ethereum, blockNumber)
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

func (app *App) StartBSC(ctx context.Context) {
	go func(ctx context.Context) {
		app.logger.Info("Starting BSC watcher...")
		for {
			select {
			case <-time.After(3 * time.Second):
				var err error
				app.bscClient, err = ethclient.Dial(app.cfg.Bsc.RpcURL)
				if err != nil {
					app.logger.Error("failed to update client", zap.Error(err))
					continue
				}

				app.swapBSC, err = web3.NewOTCSwap(common.HexToAddress(app.cfg.Bsc.SwapAddress), app.bscClient)
				if err != nil {
					app.logger.Error("failed to generate nft abi", zap.Error(err))
					continue
				}

				lastBlockNumber, err := app.blockStore.GetLastBlock(ctx, types.BSC)
				if err != nil {
					app.logger.Error("failed to get last block number", zap.Error(err))
					continue
				}
				blockNumber, err := app.bscClient.BlockNumber(ctx)
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
				} else {

				}
				for {
					app.logger.Info("getting new events...", zap.Uint64("start", start), zap.Uint64("end", end),
						zap.String("Network", fmt.Sprint(types.BSC)))
					opts := &bind.FilterOpts{
						Start:   start,
						End:     &end,
						Context: ctx,
					}
					transfers, err := app.swapBSC.FilterNewTransfer(opts)
					if err != nil {
						app.logger.Error("failed to filter swaps transfers", zap.Error(err))
						continue
					}
					for transfers.Next() {
						err := app.SyncLock(ctx, transfers.Event, types.BSC)
						if err != nil {
							app.logger.Error("failed to sync trade", zap.Error(err), zap.String("tradehash", common.BytesToHash(transfers.Event.TradeHash[:]).String()))
							continue
						}
						app.logger.Info("event catched", zap.String("Network", fmt.Sprint(types.BSC)),
							zap.String("Tradehash", common.BytesToHash(transfers.Event.TradeHash[:]).String()))
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

				_, err = app.blockStore.Update(ctx, types.BSC, blockNumber)
				if err != nil {
					app.logger.Error("failed to update last block number", zap.Error(err))
					continue
				}

				app.bscClient.Close()

			case <-ctx.Done():
				return
			}
		}
	}(ctx)
}
