package services

import (
	"auth_service/internal/repos"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type authService struct {
	authRepo repos.AuthRepo
	logger   *loggers.LoggerZap
}

func (as *authService) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	resourceIds := make([]string, 0)
	reourceName := make([]string, 0)

	actionIds := make([]string, 0)
	actionName := make([]string, 0)
	actionResourceIds := make([]string, 0)
	items := req.Items

	if len(items) == 0 {
		as.logger.ErrorString("no items to save", zap.Error(fmt.Errorf("error arise at SaveRouteResource/auth.service.go")))
		return nil, fmt.Errorf("no items to save")
	}

	// get slice
	for _, item := range items {
		if item.Resource.Id != "" {
			resourceIds = append(resourceIds, item.Resource.Id)
			reourceName = append(reourceName, item.Resource.Name)

			for _, action := range item.Actions {
				if action.Id != "" {
					actionIds = append(actionIds, action.Id)
					actionName = append(actionName, action.Name)
					actionResourceIds = append(actionResourceIds, item.Resource.Id)
				}
			}
		}
	}

	// Save resources
	err := as.authRepo.SyncResources(ctx, resourceIds, reourceName)
	if err != nil {
		as.logger.ErrorString("Failed to sync resources", zap.Error(err))
		return nil, err
	}

	err = as.authRepo.SyncActions(ctx, actionIds, actionResourceIds, actionName)
	if err != nil {
		as.logger.ErrorString("Failed to sync actions", zap.Error(err))
		return nil, err
	}

	resp := &auth.SaveRouteResourceResponse{
		Success: true,
		Message: "Resources and actions saved successfully",
	}

	return resp, nil
}

func (as *authService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// TODO: Implement login logic
	return &auth.LoginResponse{}, nil
}

func (as *authService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// TODO: Implement register logic
	return &auth.RegisterResponse{}, nil
}

func (as *authService) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error) {
	// TODO: Implement confirm email logic
	return &auth.ConfirmEmailResponse{}, nil
}

func (as *authService) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	// TODO: Implement logout logic
	return &auth.LogoutResponse{}, nil
}

func (as *authService) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
	// TODO: Implement forgot password logic
	return &auth.ForgotPasswordResponse{}, nil
}

func (as *authService) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error) {
	// TODO: Implement confirm forgot password logic
	return &auth.ConfirmForgotPasswordResponse{}, nil
}

func (as *authService) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	// TODO: Implement reset password logic
	return &auth.ResetPasswordResponse{}, nil
}
