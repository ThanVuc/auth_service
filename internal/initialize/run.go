package initialize

import (
	"auth_service/global"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Run initializes the application by loading the configuration,
establishing database connections, and setting up the HTTP server with the specified routes.
@Note: This function is the entry point for the application, setting up the necessary components
*/
func Run() {
	print("gRPC servers are running...")
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	defer cancel()

	LoadConfig()
	InitLogger()
	InitPostgreSQL()
	InitRedis()

	RunMigrations(global.PostgresPool)

	logger := global.Logger

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	NewAuthService().RunServers(ctx, wg)

	<-stop
	cancel()
	err := global.RedisDb.Close(wg)
	if err != nil {
		logger.Error("Failed to close Redis connection", "", zap.Error(err))
	} else {
		global.Logger.Info("Redis connection closed successfully", "")
	}
	if err := global.Logger.Sync(wg); err != nil {
		logger.Error("Failed to sync logger", "", zap.Error(err))
	} else {
		global.Logger.Info("Logger synced successfully", "")
	}

	wg.Wait()
}
