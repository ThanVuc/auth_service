package initialize

import (
	"auth_service/global"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func runMigrations(db *pgxpool.Pool) {
	// Folder where your migration files (.sql) are stored
	migrationsDir := "./sql/schema"
	logger := global.Logger

	// Convert *pgxpool.Pool to *sql.DB using stdlib
	sqlDB := stdlib.OpenDBFromPool(global.PostgresPool)

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		logger.Error("Failed to apply migrations:", "", zap.Error(err))
		return
	}

	logger.Info("Migrations applied successfully from", "")
}
