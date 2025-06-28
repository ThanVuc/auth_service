package repos

import (
	"auth_service/internal/database"
	"auth_service/pkg/loggers"
	"context"
)

type permissionRepo struct {
	logger *loggers.LoggerZap
	sqlc   *database.Queries
}

func (r *permissionRepo) GetResources(ctx context.Context) ([]database.GetResourcesRow, error) {
	resources, err := r.sqlc.GetResources(ctx)
	if err != nil {
		r.logger.ErrorString("failed to get resources in database")
		return nil, err
	}
	return resources, nil
}

func (r *permissionRepo) GetActions(ctx context.Context, resourceId string) ([]database.GetActionsRow, error) {
	actions, err := r.sqlc.GetActions(ctx, resourceId)
	if err != nil {
		r.logger.ErrorString("failed to get actions in database")
		return nil, err
	}
	return actions, nil
}
