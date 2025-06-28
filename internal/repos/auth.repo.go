package repos

import (
	"auth_service/internal/database"
	"auth_service/pkg/loggers"
	"context"
)

type authRepo struct {
	sqlc   *database.Queries
	logger *loggers.LoggerZap
}

// All the below methods are for testing purposes only
func (ar *authRepo) SyncResources(ctx context.Context, ids []string, names []string) error {
	err := ar.sqlc.UpsertResources(ctx, database.UpsertResourcesParams{
		Column1: ids,
		Column2: names,
	})

	if err != nil {
		ar.logger.ErrorString("Failed to upsert resource at SyncResources in auth.repo")
		return err
	}

	err = ar.sqlc.RemoveOldResources(ctx, ids)
	if err != nil {
		ar.logger.ErrorString("Failed to delete resources not in use at SyncResources in auth.repo")
		return err
	}

	return nil
}

func (ar *authRepo) SyncActions(ctx context.Context, ids, resourceIds, names []string) error {
	err := ar.sqlc.UpsertActions(ctx, database.UpsertActionsParams{
		Column1: ids,
		Column2: resourceIds,
		Column3: names,
	})

	if err != nil {
		ar.logger.ErrorString("Failed to upsert action at SyncActions in auth.repo")
		return err
	}

	err = ar.sqlc.RemoveOldActions(ctx, ids)
	if err != nil {
		ar.logger.ErrorString("Failed to delete actions not in use at SyncActions in auth.repo")
		return err
	}

	return nil
}
