package app

import (
	"context"

	"go.uber.org/zap"
)

func (app *App) SyncTrade(ctx context.Context, tradehash string) error {
	trade, err := app.GetTrade(tradehash)
	if err != nil {
		app.logger.Warn("failed to get trade by tradehash", zap.Error(err), zap.String("tradehash", tradehash))
		return err
	}
	err = app.UpdateTrade(ctx, trade)
	if err != nil {
		app.logger.Warn("failed to update trade", zap.Error(err), zap.Reflect("trade", trade))
		return err
	}
	return nil
}
