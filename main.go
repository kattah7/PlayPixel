package main

import (
	"context"
	"fmt"
	"os"

	"github.com/PlayPixel/api/internal/api"
	"github.com/PlayPixel/api/internal/db"
	"github.com/PlayPixel/api/internal/logger"
	"github.com/PlayPixel/api/pkg/config"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()

	var atomicLogLevel zap.AtomicLevel
	var err error

	if atomicLogLevel, err = zap.ParseAtomicLevel(cfg.LogLevel); err != nil {
		fmt.Printf("Invalid log level supplied (%s), defaulting to info\n", err)
		atomicLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	log := logger.New(atomicLogLevel, cfg.LogDevelopment)
	defer log.Sync()

	ctx := logger.OnContext(context.Background(), log)

	pool, err := db.NewPool(ctx, cfg.DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		log.Fatal(err)
	}

	if err := db.Run(ctx, pool); err != nil {
		fmt.Fprintf(os.Stderr, "unable to run migrations: %v\n", err)
		log.Fatal(err)
	}

	s := api.Init(ctx, *cfg, pool, log)
	s.Run()
}
