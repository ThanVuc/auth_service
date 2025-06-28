package services

import (
	"auth_service/global"
	"auth_service/internal/grpc/auth"
	"auth_service/internal/repos"
	"context"
)

type (
	AuthService interface {
		SaveRouteResource(ctx context.Context, items *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
	}

	PermissionService interface {
		GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error)
	}

	RoleService interface {
	}

	TokenService interface{}
)

func NewAuthService(
	authRepo repos.AuthRepo,
) AuthService {
	return &authService{
		authRepo: authRepo,
		logger:   global.Logger,
	}
}

func NewPermissionService(permissionRepo repos.PermissionRepo) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func NewRoleService(roleRepo repos.RoleRepo) RoleService {
	return &roleService{
		repo: roleRepo,
	}
}

func NewTokenService(tokenRepo repos.TokenRepo) TokenService {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}
