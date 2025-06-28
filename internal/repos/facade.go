package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
	"context"
)

type (
	AuthRepo interface {
		SyncResources(ctx context.Context, ids []string, names []string) error
		SyncActions(ctx context.Context, ids, resourceIds, names []string) error
	}

	PermissionRepo interface {
		GetResources(ctx context.Context) ([]database.GetResourcesRow, error)
		GetActions(ctx context.Context, resourceId string) ([]database.GetActionsRow, error)
	}

	RoleRepo interface{}

	TokenRepo interface{}
)

func NewAuthRepo() AuthRepo {
	return &authRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewPermissionRepo() PermissionRepo {
	return &permissionRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewRoleRepo() RoleRepo {
	return &roleRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

func NewTokenRepo() TokenRepo {
	return &tokenRepo{
		redisDb: global.RedisDb,
		logger:  global.Logger,
	}
}
