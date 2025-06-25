package repos

import (
	"auth_service/global"
	"auth_service/internal/database"
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/pkg/loggers"
	"context"

	"go.uber.org/zap"
)

type IAuthRepo interface {
	UpsertResource(ctx context.Context, resource *v1.Resource) error
	UpsertAction(ctx context.Context, resourceId string, action *v1.Action) error
	DeleteResourceNotInUse(ctx context.Context, resourceIds []string) error
	DeleteActionNotInUse(ctx context.Context, actionIds []string) error
}

type AuthRepo struct {
	sqlc   *database.Queries
	logger *loggers.LoggerZap
}

func NewAuthRepo() IAuthRepo {
	return &AuthRepo{
		sqlc:   database.New(global.PostgresPool),
		logger: global.Logger,
	}
}

// All the below methods are for testing purposes only
func (ar *AuthRepo) UpsertResource(ctx context.Context, resource *v1.Resource) error {
	err := ar.sqlc.UpsertResourceByID(ctx, database.UpsertResourceByIDParams{
		ResourceID: resource.ResourceId,
		Name:       resource.Resource,
	})
	if err != nil {
		ar.logger.ErrorString("Failed to upsert resource", zap.Error(err))
		return err
	}

	return nil
}

func (ar *AuthRepo) UpsertAction(ctx context.Context, resourceId string, action *v1.Action) error {
	err := ar.sqlc.UpsertActionByID(ctx, database.UpsertActionByIDParams{
		ActionID:   action.ActionId,
		Name:       action.Action,
		ResourceID: resourceId,
	})
	if err != nil {
		ar.logger.ErrorString("Failed to upsert action", zap.Error(err))
		return err
	}

	return nil
}

func (ar *AuthRepo) DeleteResourceNotInUse(ctx context.Context, resourceIds []string) error {
	if err := ar.sqlc.DeleteResourceNotInUse(ctx, resourceIds); err != nil {
		ar.logger.ErrorString("Failed to delete resource not in use", zap.Error(err))
		return err
	}

	return nil
}

func (ar *AuthRepo) DeleteActionNotInUse(ctx context.Context, actionIds []string) error {
	if err := ar.sqlc.DeleteActionNotInUse(ctx, actionIds); err != nil {
		ar.logger.ErrorString("Failed to delete action not in use", zap.Error(err))
		return err
	}

	return nil
}
