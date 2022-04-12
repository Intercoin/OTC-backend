package main

import (
	"os"
	"os/signal"
	"otc-backend/app"
	"otc-backend/repository/postgres"
	"syscall"

	"go.uber.org/zap"

	"otc-backend/config"
)

func main() {
	cfg := config.Get()

	logger, _ := zap.NewDevelopment()
	defer func() {
		_ = logger.Sync()
	}()

	db, err := postgres.Connect(cfg.DB)
	if err != nil {
		logger.Fatal("failed to connect db", zap.Error(err))
	}

	defer func() {
		_ = db.Close()
	}()

	swapRepo := postgres.NewSwapRepository(db)

	_, err = app.NewApp(logger, cfg, swapRepo)
	if err != nil {
		logger.Fatal("failed to create app", zap.Error(err))
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
