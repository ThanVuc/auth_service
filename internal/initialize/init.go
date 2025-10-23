package initialize

import (
	"auth_service/global"
	"context"
	"sync"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

func initConfigAndResources() error {
	loadConfig()
	initLogger()
	initPostgreSQL()
	initRedis()
	initEventBus()
	runMigrations(global.PostgresPool)
	initR2()

	return nil
}

func startGrpcSerivces(ctx context.Context, wg *sync.WaitGroup) {
	authService := NewAuthService()
	authService.runServers(ctx, wg)
}

func gracefulShutdown(wg *sync.WaitGroup, logger log.Logger) {
	wg.Add(1)
	err := global.RedisDb.Close(wg)
	handleError(logger, err, "Redis connection closed successfully")

	wg.Add(1)
	err = logger.Sync(wg)
	handleError(logger, err, "Logger synced successfully")

	global.PostgresPool.Close()

	wg.Add(1)
	global.EventBusConnector.Close(wg)

	wg.Wait()
}

func handleError(logger log.Logger, err error, successMessage string) {
	if err != nil {
		logger.Error("An error occurred", "", zap.Error(err))
	} else {
		logger.Info(successMessage, "")
	}
}
