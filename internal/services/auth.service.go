package services

import (
	"auth_service/global"
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/internal/repos"
	"auth_service/pkg/loggers"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type IAuthService interface {
	SaveRouteResource(ctx context.Context, items []*v1.ResourceItem) bool
}

type AuthService struct {
	authRepo repos.IAuthRepo
	logger   *loggers.LoggerZap
}

func NewAuthService(
	authRepo repos.IAuthRepo,
) IAuthService {
	return &AuthService{
		authRepo: authRepo,
		logger:   global.Logger,
	}
}

func (as *AuthService) SaveRouteResource(ctx context.Context, items []*v1.ResourceItem) bool {
	resourceIds := make([]string, 0)
	reourceName := make([]string, 0)

	actionIds := make([]string, 0)
	actionName := make([]string, 0)
	actionResourceIds := make([]string, 0)

	if len(items) == 0 {
		as.logger.ErrorString("no items to save", zap.Error(fmt.Errorf("error arise at SaveRouteResource/auth.service.go")))
		return false
	}

	// get slice
	for _, item := range items {
		if item.Resource.ResourceId != "" {
			resourceIds = append(resourceIds, item.Resource.ResourceId)
			reourceName = append(reourceName, item.Resource.Resource)

			for _, action := range item.Actions {
				if action.ActionId != "" {
					actionIds = append(actionIds, action.ActionId)
					actionName = append(actionName, action.Action)
					actionResourceIds = append(actionResourceIds, item.Resource.ResourceId)
				}
			}
		}
	}

	// Save resources
	err := as.authRepo.SyncResources(ctx, resourceIds, reourceName)
	if err != nil {
		as.logger.ErrorString("Failed to sync resources", zap.Error(err))
		return false
	}

	err = as.authRepo.SyncActions(ctx, actionIds, actionResourceIds, actionName)
	if err != nil {
		as.logger.ErrorString("Failed to sync actions", zap.Error(err))
		return false
	}

	return true
}
