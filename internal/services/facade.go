package services

import (
	"auth_service/global"
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
)

type (
	AuthService interface {
		SaveRouteResource(ctx context.Context, items *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
		LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error)
		Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error)
	}

	PermissionService interface {
		GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error)
		GetActions(ctx context.Context, items *auth.GetActionsRequest) (*auth.GetActionsResponse, error)
		GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error)
		GetPermission(ctx context.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error)
		UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error)
		DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error)
	}

	RoleService interface {
		GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error)
		UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error)
		DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error)
		DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error)
		GetRoleById(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error)
	}

	TokenService interface {
		RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
		RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error)
	}
)

func NewAuthService(
	authRepo repos.AuthRepo,
) AuthService {
	return &authService{
		authRepo: authRepo,
		logger:   global.Logger,
	}
}

func NewPermissionService(
	permissionRepo repos.PermissionRepo,
	mapper mapper.PermissionMapper,
) PermissionService {
	return &permissionService{
		permissionRepo:   permissionRepo,
		permissionMapper: mapper,
		logger:           global.Logger,
		pool:             global.PostgresPool,
	}
}

func NewRoleService(
	roleRepo repos.RoleRepo,
	roleMapper mapper.RoleMapper,
) RoleService {
	return &roleService{
		roleRepo:   roleRepo,
		roleMapper: roleMapper,
		pool:       global.PostgresPool,
	}
}

func NewTokenService(tokenRepo repos.TokenRepo) TokenService {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}
