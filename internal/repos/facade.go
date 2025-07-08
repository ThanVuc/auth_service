package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
	"auth_service/proto/auth"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	AuthRepo interface {
		SyncResources(ctx context.Context, ids []string, names []string) error
		SyncActions(ctx context.Context, ids, resourceIds, names []string) error
		Login(ctx context.Context, req *auth.LoginRequest) error
		Register(ctx context.Context, req *auth.RegisterRequest) error
		ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) error
		Logout(ctx context.Context, req *auth.LogoutRequest) error
		ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) error
		ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) error
		ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) error
	}

	PermissionRepo interface {
		GetResources(ctx context.Context) ([]database.GetResourcesRow, error)
		GetActions(ctx context.Context, resourceId string) ([]database.GetActionsRow, error)
		GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) ([]database.GetPermissionsRow, int32, int32, error)
		UpsertPermission(ctx context.Context, tx pgx.Tx, req *auth.UpsertPermissionRequest) (*pgtype.UUID, error)
		GetActionsByPermissionId(ctx context.Context, tx pgx.Tx, permId pgtype.UUID) ([]string, error)
		UpdateActionsToPermission(ctx context.Context, tx pgx.Tx, permId pgtype.UUID, addActionIds []string, deleteActionIds []string) error
		DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (bool, error)
		GetPermission(ctx context.Context, req *auth.GetPermissionRequest) (*[]database.GetPermissionRow, error)
	}

	RoleRepo interface {
		GetRoles(ctx context.Context, req *auth.GetRolesRequest) ([]database.GetRolesRow, int32, int32, error)
		GetRoleById(ctx context.Context, req *auth.GetRoleRequest) (*[]database.GetRoleByIdRow, error)
		UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) error
		DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (bool, error)
		DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (bool, error)
		CountUsersByRoles(ctx context.Context, roleIds []pgtype.UUID) (*[]database.CountUsersByRolesRow, error)
	}

	TokenRepo interface {
		RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) error
		RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) error
	}
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
