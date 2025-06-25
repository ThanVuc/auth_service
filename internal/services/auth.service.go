package services

import (
	"auth_service/global"
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/internal/repos"
	"auth_service/pkg/loggers"
	"context"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IAuthService interface {
	SaveRouteResource(items []*v1.ResourceItem) bool
}

type AuthService struct {
	authRepo repos.IAuthRepo
	pgx      *pgxpool.Pool
	logger   *loggers.LoggerZap
}

func NewAuthService(
	authRepo repos.IAuthRepo,
) IAuthService {
	return &AuthService{
		authRepo: authRepo,
		pgx:      global.PostgresPool,
		logger:   global.Logger,
	}
}

func (as *AuthService) SaveRouteResource(items []*v1.ResourceItem) bool {
	ctx := context.Background()
	tx, err := as.pgx.Begin(ctx)
	resourceIds := make([]string, 0)
	actionIds := make([]string, 0)

	if err != nil {
		as.logger.ErrorString("Failed to begin transaction", zap.Error(err))
		return false
	}

	for _, item := range items {
		resource := item.Resource
		resourceIds = append(resourceIds, resource.ResourceId)
		err := as.authRepo.UpsertResource(ctx, resource)
		if err != nil {
			as.logger.ErrorString("Failed to upsert resource", zap.Error(err))
			tx.Rollback(ctx)
			return false
		}

		for _, action := range item.Actions {
			actionIds = append(actionIds, action.ActionId)
			err := as.authRepo.UpsertAction(ctx, resource.ResourceId, action)
			if err != nil {
				as.logger.ErrorString("Failed to upsert action", zap.Error(err))
				tx.Rollback(ctx)
				return false
			}
		}
	}

	if err := as.authRepo.DeleteResourceNotInUse(ctx, resourceIds); err != nil {
		as.logger.ErrorString("Failed to delete resource not in use", zap.Error(err))
		tx.Rollback(ctx)
		return false
	}
	if err := as.authRepo.DeleteActionNotInUse(ctx, actionIds); err != nil {
		as.logger.ErrorString("Failed to delete action not in use", zap.Error(err))
		tx.Rollback(ctx)
		return false
	}

	tx.Commit(ctx)

	return true
}
