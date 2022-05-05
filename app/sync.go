package app

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"otc-backend/types"

	"otc-backend/web3"
)

func (app *App) SyncLock(ctx context.Context, event *web3.OTCSwapNewTransfer, network types.Network) error {

	lock := event.ToModel()
	lock.Network = network

	trade, err := app.GetTrade(ctx, lock.TradeHash)
	if err != nil {
		if err == sql.ErrNoRows {
			trade = event.ToTradeModel(network)
			err = app.CreateTrade(ctx, trade)
			if err != nil {
				app.logger.Warn("failed to create new trade", zap.Error(err), zap.Reflect("trade", trade))
				return err
			}
		} else {
			app.logger.Warn("failed to get trade", zap.Error(err), zap.Reflect("trade", trade))
			return err
		}
	}

	if trade != nil && trade.Network1 != network && trade.Network2 == types.Unknown && (trade.Address1 == lock.Address || trade.Address2 == lock.Address) {
		err = app.UpdateTrade(ctx, trade.ID, network)
		if err != nil {
			app.logger.Warn("failed to update trade", zap.Error(err), zap.Reflect("trade", trade))
			return err
		}
	}

	err = app.SetLock(ctx, lock)
	if err != nil {
		app.logger.Warn("failed to write new lock", zap.Error(err), zap.Reflect("lock", lock))
		return err
	}
	return nil
}

func (app *App) SyncEngage(ctx context.Context, event *web3.OTCSwapEngaged, network types.Network) error {

	engage := event.ToModel()
	engage.Network = network

	err := app.SetEngage(ctx, engage)
	if err != nil {
		app.logger.Warn("failed to write new engage", zap.Error(err), zap.Reflect("engage", engage))
		return err
	}
	return nil
}

func (app *App) SyncClaim(ctx context.Context, event *web3.OTCSwapClaimed, network types.Network) error {

	claim := event.ToModel()
	claim.Network = network

	err := app.SetClaim(ctx, claim)
	if err != nil {
		app.logger.Warn("failed to write new claim", zap.Error(err), zap.Reflect("claim", claim))
		return err
	}
	return nil
}
