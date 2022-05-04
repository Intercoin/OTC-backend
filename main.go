package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"otc-backend/app"
	"otc-backend/config"
	"otc-backend/repository/postgres"
	"otc-backend/server"
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
	blockRepo := postgres.NewBlockRepository(db)

	app, err := app.NewApp(logger, cfg, swapRepo, blockRepo)
	if err != nil {
		logger.Fatal("failed to create app", zap.Error(err))
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	app.Start(ctx)
	app.StartBSC(ctx)

	logger.Info("server is starting...")
	srv := server.NewServer(logger, cfg.Addr, app)
	srv.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
