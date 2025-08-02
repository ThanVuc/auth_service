package repos

import (
	"auth_service/internal/database"
	"context"

	"github.com/thanvuc/go-core-lib/log"
)

type authRepo struct {
	sqlc   *database.Queries
	logger log.Logger
}

// All the below methods are for testing purposes only
func (ar *authRepo) SyncResources(ctx context.Context, ids []string, names []string) error {
	err := ar.sqlc.UpsertResources(ctx, database.UpsertResourcesParams{
		Column1: ids,
		Column2: names,
	})

	if err != nil {
		return err
	}

	err = ar.sqlc.RemoveOldResources(ctx, ids)
	if err != nil {
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
		return err
	}

	err = ar.sqlc.RemoveOldActions(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
