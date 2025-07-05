package repos

import (
	"auth_service/internal/database"
	"auth_service/internal/utils"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type permissionRepo struct {
	logger *loggers.LoggerZap
	sqlc   *database.Queries
	pool   *pgxpool.Pool
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

func (r *permissionRepo) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) ([]database.GetPermissionsRow, int32, int32, error) {
	pagination := utils.ToPagination(req.PageQuery.Page, req.PageQuery.PageSize)
	permissions, err := r.sqlc.GetPermissions(
		ctx,
		database.GetPermissionsParams{
			Column1: req.Search,
			Column2: req.ResourceId,
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
		},
	)
	if err != nil {
		r.logger.ErrorString("failed to get permissions in database")
		return nil, 0, 0, err
	}

	total_perms, err := r.sqlc.CountRootPermissions(
		ctx,
		database.CountRootPermissionsParams{
			Column1: req.Search,
			Column2: req.ResourceId,
		},
	)

	if err != nil {
		r.logger.ErrorString("failed to count permissions in database")
		return nil, 0, 0, err
	}

	total_roots, err := r.sqlc.CountRootPermissions(
		ctx,
		database.CountRootPermissionsParams{
			Column1: req.Search,
			Column2: req.ResourceId,
		},
	)

	return permissions, int32(total_perms), int32(total_roots), err
}

func (r *permissionRepo) UpsertPermission(ctx context.Context, tx pgx.Tx, req *auth.UpsertPermissionRequest) (*pgtype.UUID, error) {
	sqlcTx := r.sqlc.WithTx(tx)

	if req.PermissionId != nil {
		// Update existing permission
		var permUUID pgtype.UUID
		if err := permUUID.Scan(*req.PermissionId); err != nil {
			r.logger.ErrorString("failed to parse permission id to UUID")
			return nil, err
		}

		permId, err := sqlcTx.UpdatePermission(
			ctx,
			database.UpdatePermissionParams{
				PermID:      permUUID,
				Name:        req.Name,
				ResourceID:  req.ResourceId,
				Description: pgtype.Text{String: req.Description, Valid: true},
			},
		)

		if err != nil && err == sql.ErrNoRows {
			r.logger.ErrorString("The permission does not exist in the database")
			return nil, err
		}

		if err != nil {
			r.logger.ErrorString("failed to update permission in database")
			return nil, err
		}

		return &permId, nil
	}
	// Create new permission
	newPermId, err := sqlcTx.InsertPermission(
		ctx,
		database.InsertPermissionParams{
			Name:        req.Name,
			ResourceID:  req.ResourceId,
			Description: pgtype.Text{String: req.Description, Valid: true},
		},
	)

	if err != nil {
		r.logger.ErrorString("failed to create permission in database")
		return nil, err
	}

	return &newPermId, nil
}

func (r *permissionRepo) GetActionsByPermissionId(ctx context.Context, tx pgx.Tx, permId pgtype.UUID) ([]string, error) {
	sqlcTx := r.sqlc.WithTx(tx)
	actions, err := sqlcTx.GetActionsByPermissionId(ctx, permId)
	if err != nil {
		r.logger.ErrorString("failed to get actions by permission id in database")
		return nil, err
	}

	actionIds := make([]string, 0, len(actions))
	for _, action := range actions {
		actionIds = append(actionIds, action.ActionID)
	}
	return actionIds, nil
}

func (r *permissionRepo) UpdateActionsToPermission(ctx context.Context, tx pgx.Tx, permId pgtype.UUID, addActionIds []string, deleteActionIds []string) error {
	sqlcTx := r.sqlc.WithTx(tx)
	if len(addActionIds) != 0 {
		err := sqlcTx.AddActionToPermission(
			ctx,
			database.AddActionToPermissionParams{
				PermID:  permId,
				Column2: addActionIds,
			},
		)
		if err != nil {
			r.logger.ErrorString("failed to add actions to permission in database")
			return err
		}
	}

	if len(deleteActionIds) != 0 {
		err := sqlcTx.DeleteActionToPermission(
			ctx,
			database.DeleteActionToPermissionParams{
				PermID:  permId,
				Column2: deleteActionIds,
			},
		)
		if err != nil {
			r.logger.ErrorString("failed to delete actions from permission in database")
			return err
		}
	}

	return nil
}

func (r *permissionRepo) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) error {
	// TODO: Implement delete permission logic
	return nil
}

func (r *permissionRepo) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) error {
	// TODO: Implement assign permission to role logic
	return nil
}
