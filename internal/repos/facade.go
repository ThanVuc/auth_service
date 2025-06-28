package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
	"context"
)

type (
	IAuthRepo interface {
		SyncResources(ctx context.Context, ids []string, names []string) error
		SyncActions(ctx context.Context, ids, resourceIds, names []string) error
	}

	IPermissionRepo interface{}

	IRoleRepo interface{}

	ITokenRepo interface{}
)

func NewAuthRepo() IAuthRepo {
	return &AuthRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewPermissionRepo() IPermissionRepo {
	return &PermissionRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewRoleRepo() IRoleRepo {
	return &RoleRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewTokenRepo() ITokenRepo {
	return &TokenRepo{
		redisDb: global.RedisDb,
		logger:  global.Logger,
	}
}
