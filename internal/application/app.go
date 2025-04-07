package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/xamust/ekom-test/config"
	v1 "github.com/xamust/ekom-test/internal/controller/http"
	"github.com/xamust/ekom-test/internal/mappers"
	"github.com/xamust/ekom-test/internal/repository"
	"github.com/xamust/ekom-test/internal/usecases"
	"github.com/xamust/ekom-test/pkg/db/mongo"
	"github.com/xamust/ekom-test/pkg/http/fiber"
	"github.com/xamust/ekom-test/pkg/logger"
)

func RunWithCtx(ctx context.Context, cfg *config.Config) {
	// logger
	log := logger.NewLogger()

	// db
	mongoDB, err := mongo.NewMongoDBConnectionWithCtx(ctx, log, cfg)
	if err != nil {
		log.Fatal(err)
	}
	db := mongoDB.Client.Database(cfg.DB.Mongo.DBName)
	repo := repository.NewRepository(log, db, cfg)

	// usecases
	deps := &usecases.Dependencies{
		Log:    log,
		Mapper: mappers.NewMappers(),
		Repo:   repo,
	}

	usecase := usecases.NewUsecases(deps)

	// httpServer
	httpServer := fiber.New(
		fiber.WithPort(cfg.HTTP.Port),
		fiber.WithPrefork(cfg.HTTP.UsePreforkMode))
	v1.NewRouter(httpServer.App, cfg, usecase, log)

	// Start servers
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
