package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
)

type IAuthRepo interface {
}

type AuthRepo struct {
	sqlc *database.Queries
}

func NewAuthRepo() IAuthRepo {
	return &AuthRepo{
		sqlc: database.New(global.PostgresPool),
	}
}

// All the below methods are for testing purposes only
