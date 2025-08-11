package services

import (
	"auth_service/internal/grpc/repos"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
	"fmt"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type authService struct {
	authRepo repos.AuthRepo
	logger   log.Logger
}

func (as *authService) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	resourceIds := make([]string, 0)
	reourceName := make([]string, 0)

	actionIds := make([]string, 0)
	actionName := make([]string, 0)
	actionResourceIds := make([]string, 0)
	items := req.Items

	if len(items) == 0 {
		as.logger.Error("no items to save: error arise at SaveRouteResource/auth.service.go", "", zap.Error(fmt.Errorf("no items to save")))
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
		as.logger.Error("Failed to sync resources", "", zap.Error(err))
		return nil, err
	}

	err = as.authRepo.SyncActions(ctx, actionIds, actionResourceIds, actionName)
	if err != nil {
		as.logger.Error("Failed to sync actions", "", zap.Error(err))
		return nil, err
	}

	resp := &auth.SaveRouteResourceResponse{
		Success: true,
		Message: "Resources and actions saved successfully",
	}

	return resp, nil
}

func (as *authService) LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error) {
	// TODO: Implement login logic
	return &auth.LoginWithGoogleResponse{}, nil
}

func (as *authService) Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error) {
	// TODO: Implement logout logic
	return &common.EmptyResponse{}, nil
}
