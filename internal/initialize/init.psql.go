package initialize

import (
	"auth_service/global"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func initPostgreSQL() {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	configs := global.Config.Postgres
	println(dsn)
	var connectString = fmt.Sprintf(dsn, configs.Host, configs.User, configs.Password, configs.Database, configs.Port)
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(connectString)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse PostgreSQL connection string: %v", err))
	}
	setPostgresConfig(config)
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create PostgreSQL connection pool: %v", err))
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		panic(fmt.Sprintf("Failed to ping PostgreSQL database: %v", err))
	}

	global.PostgresPool = pool
}

func setPostgresConfig(config *pgxpool.Config) {
	postConfig := global.Config.Postgres
	config.MaxConns = int32(postConfig.MaxOpenConns)
	config.MaxConnIdleTime = time.Duration(postConfig.ConnMaxIdleTime) * time.Second
	config.MaxConnLifetime = time.Duration(postConfig.MaxLifetime) * time.Second
}
