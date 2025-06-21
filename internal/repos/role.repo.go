package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
)

type IRoleRepo interface {
}

type RoleRepo struct {
	sqlc *database.Queries
}

func NewRoleRepo() IRoleRepo {
	return &RoleRepo{
		sqlc: database.New(global.PostgresPool),
	}
}
