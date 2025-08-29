package services

import (
	"auth_service/global"
	"auth_service/internal/grpc/helper"
	"auth_service/internal/grpc/mapper"
	"auth_service/internal/grpc/repos"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
)

type (
	AuthService interface {
		SaveRouteResource(ctx context.Context, items *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
		LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error)
		Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error)
		RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
		CheckPermission(ctx context.Context, req *auth.CheckPermissionRequest) (*auth.CheckPermissionResponse, error)
		GetUserActionsAndResources(ctx context.Context, req *auth.GetUserActionsAndResourcesRequest) (*auth.GetUserActionsAndResourcesResponse, error)
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

	UserService interface {
		GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error)
		AssignRoleToUser(ctx context.Context, req *auth.AssignRoleToUserRequest) (*common.EmptyResponse, error)
	}
)

func NewAuthService(
	authRepo repos.AuthRepo,
	jwtHelper helper.JWTHelper,
	authMapper mapper.AuthMapper,
) AuthService {
	return &authService{
		authRepo:   authRepo,
		logger:     global.Logger,
		jwtHelper:  jwtHelper,
		authMapper: authMapper,
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

func NewUserService(
	userRepo repos.UserRepo,
	userMapper mapper.UserMapper,
) UserService {
	return &userService{
		userRepo:   userRepo,
		userMapper: userMapper,
		logger:     global.Logger,
	}
}
