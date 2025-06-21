package initialize

import (
	"auth_service/global"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func checkInitError(err error) {
	if err != nil {
		panic("failed to initialize PostgreSQL: " + err.Error())
	}
}

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Initialize the PostgreSQL database connection and set up the connection pool.
This function uses the pgxpool package to create a connection pool for PostgreSQL.
It constructs the connection string using the configuration from global.Config.Postgres.
@Note: The PostgreSQL connection pool is stored in global.PostgresPool for use throughout the application.
*/
func InitPostgreSQL() {
	logger := global.Logger
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	configs := global.Config.Postgres
	var connectString = fmt.Sprintf(dsn, configs.Host, configs.User, configs.Password, configs.Database, configs.Port)
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(connectString)
	checkInitError(err)
	setPostgresConfig(config)
	pool, err := pgxpool.NewWithConfig(ctx, config)
	checkInitError(err)
	global.PostgresPool = pool
	logger.InfoString("PostgreSQL connection pool initialized")
}

func setPostgresConfig(config *pgxpool.Config) {
	postConfig := global.Config.Postgres
	config.MaxConns = int32(postConfig.MaxOpenConns)
	config.MaxConnIdleTime = time.Duration(postConfig.ConnMaxIdleTime) * time.Second
	config.MaxConnLifetime = time.Duration(postConfig.MaxLifetime) * time.Second
}
