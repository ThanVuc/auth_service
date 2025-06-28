package repos

import (
	"auth_service/internal/database"
	"auth_service/pkg/loggers"
)

type PermissionRepo struct {
	logger *loggers.LoggerZap
	sqlc   *database.Queries
}
