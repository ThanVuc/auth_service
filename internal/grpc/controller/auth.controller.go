package controller

import (
	"auth_service/internal/grpc/helper"
	"auth_service/internal/grpc/services"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
)

type AuthController struct {
	auth.UnimplementedAuthServiceServer
	authService services.AuthService
	common.UnimplementedSyncDatabaseServiceServer
}

func NewAuthController(authService services.AuthService, helper helper.JWTHelper) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.LoginWithGoogle)
}

func (ac *AuthController) Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.Logout)
}

func (ac *AuthController) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.SaveRouteResource)
}

func (ac *AuthController) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.RefreshToken)
}

func (ac *AuthController) CheckPermission(ctx context.Context, req *auth.CheckPermissionRequest) (*auth.CheckPermissionResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.CheckPermission)
}

func (ac *AuthController) GetUserActionsAndResources(ctx context.Context, req *auth.GetUserActionsAndResourcesRequest) (*auth.GetUserActionsAndResourcesResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.GetUserActionsAndResources)
}

func (ac *AuthController) SyncDatabase(ctx context.Context, req *common.SyncDatabaseRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.SyncDatabase)
}
