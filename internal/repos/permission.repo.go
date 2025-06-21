package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
)

type IPermissionRepo interface{}

type PermissionRepo struct {
	sqlc *database.Queries
}

func NewPermissionRepo() IPermissionRepo {
	return &PermissionRepo{
		sqlc: database.New(global.PostgresPool),
	}
}
