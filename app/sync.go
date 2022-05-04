package app

import (
	"context"

	"otc-backend/web3"
)

func (app *App) SyncLock(ctx context.Context, event *web3.OTCSwapNewTransfer) error {

	//err := app.UpdateTrade(ctx, trade)
	//if err != nil {
	//	app.logger.Warn("failed to a", zap.Error(err), zap.Reflect("trade", trade))
	//	return err
	//}
	return nil
}
