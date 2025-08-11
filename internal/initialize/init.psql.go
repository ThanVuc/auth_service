package initialize

import (
	"auth_service/global"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Initialize the PostgreSQL database connection and set up the connection pool.
This function uses the pgxpool package to create a connection pool for PostgreSQL.
It constructs the connection string using the configuration from global.Config.Postgres.
@Note: The PostgreSQL connection pool is stored in global.PostgresPool for use throughout the application.
*/
func initPostgreSQL() {
	logger := global.Logger
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	configs := global.Config.Postgres
	println(dsn)
	var connectString = fmt.Sprintf(dsn, configs.Host, configs.User, configs.Password, configs.Database, configs.Port)
	ctx := context.Background()
	for {
		config, err := pgxpool.ParseConfig(connectString)
		if err != nil {
			logger.Error("Failed to parse PostgreSQL connection string", "", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}
		setPostgresConfig(config)
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			logger.Error("Failed to create PostgreSQL connection pool", "", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}

		if err := pool.Ping(ctx); err != nil {
			logger.Error("Failed to ping PostgreSQL", "", zap.Error(err))
			pool.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		global.PostgresPool = pool
		break
	}
	logger.Info("PostgreSQL connection pool initialized", "")
}

func setPostgresConfig(config *pgxpool.Config) {
	postConfig := global.Config.Postgres
	config.MaxConns = int32(postConfig.MaxOpenConns)
	config.MaxConnIdleTime = time.Duration(postConfig.ConnMaxIdleTime) * time.Second
	config.MaxConnLifetime = time.Duration(postConfig.MaxLifetime) * time.Second
}
