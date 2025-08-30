package repos

import (
	"auth_service/global"
	"auth_service/internal/constant"
	"auth_service/internal/grpc/database"
	"auth_service/internal/grpc/models"
	"auth_service/proto/auth"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	AuthRepo interface {
		SyncResources(ctx context.Context, ids []string, names []string) error
		SyncActions(ctx context.Context, ids, resourceIds, names []string) error
		RegisterUserWithExternalProvider(ctx context.Context, userInfo models.GoogleUserInfo, provider constant.Provider) (string, string, error)
		LoginWithExternalProvider(ctx context.Context, sub string, email string) (*database.LoginWithExternalProviderRow, []pgtype.UUID, error)
		CheckPermission(ctx context.Context, roleIDs []string, resource string, action string) (bool, error)
		GetUserActionsAndResources(ctx context.Context, roleIDs []string) ([]database.GetUserAuthInfoRow, error)
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
		UpsertRole(ctx context.Context, tx pgx.Tx, req *auth.UpsertRoleRequest) (string, error)
		DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (bool, error)
		DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (bool, error)
		CountUsersByRoles(ctx context.Context, roleIds []pgtype.UUID) (*[]database.CountUsersByRolesRow, error)
		GetPermissionIdsByRole(ctx context.Context, tx pgx.Tx, roleId string) ([]pgtype.UUID, error)
		UpsertPermissionsForRole(ctx context.Context, tx pgx.Tx, roleId string, addPerms *[]pgtype.UUID, delPerms *[]pgtype.UUID) (bool, error)
	}

	UserRepo interface {
		GetUsers(ctx context.Context, req *auth.GetUsersRequest) ([]database.GetUsersRow, int32, int32, error)
		GetRoleIDsByUserID(ctx context.Context, userId pgtype.UUID) ([]pgtype.UUID, error)
		AddNewRolesToUser(ctx context.Context, tx pgx.Tx, userId pgtype.UUID, ids []pgtype.UUID) error
		RemoveRolesFromUser(ctx context.Context, tx pgx.Tx, userId pgtype.UUID, ids []pgtype.UUID) error
		AssignRoleToUser(ctx context.Context, req *auth.AssignRoleToUserRequest) error
		GetUser(ctx context.Context, req *auth.GetUserRequest) (*[]database.GetUserRow, error)
		LockOrUnLockUser(ctx context.Context, req *auth.LockUserRequest) (bool , error)
	}
)

func NewAuthRepo() AuthRepo {
	return &authRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
		pool:   global.PostgresPool,
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

func NewUserRepo() UserRepo {
	return &userRepo{
		logger: global.Logger,
		sqlc:   database.New(global.PostgresPool),
		pool:   global.PostgresPool,
	}
}
