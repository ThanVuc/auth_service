package repos

import (
	"auth_service/internal/database"
	"auth_service/internal/utils"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type roleRepo struct {
	logger *loggers.LoggerZap
	sqlc   *database.Queries
}

func (r *roleRepo) GetRoles(ctx context.Context, req *auth.GetRolesRequest) ([]database.GetRolesRow, int32, int32, error) {
	pagination := utils.ToPagination(req.PageQuery.Page, req.PageQuery.PageSize)

	roles, err := r.sqlc.GetRoles(ctx, database.GetRolesParams{
		Column1: req.Search,
		Limit:   pagination.Limit,
		Offset:  pagination.Offset,
	})

	if err != nil {
		return nil, 0, 0, err
	}

	totalRoles, err := r.sqlc.CountTotalRoles(ctx, req.Search)
	if err != nil {
		return nil, 0, 0, err
	}

	rootRoles, err := r.sqlc.CountRootRoles(ctx, req.Search)
	if err != nil {
		return nil, 0, 0, err
	}

	return roles, int32(totalRoles), int32(rootRoles), nil
}

func (r *roleRepo) CountUsersByRoles(ctx context.Context, roleIds []pgtype.UUID) (*[]database.CountUsersByRolesRow, error) {
	roles, err := r.sqlc.CountUsersByRoles(ctx, roleIds)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (r *roleRepo) GetRoleById(ctx context.Context, req *auth.GetRoleRequest) (*[]database.GetRoleByIdRow, error) {
	roleId, err := utils.ToUUID(req.RoleId)
	if err != nil {
		return nil, err
	}

	role, err := r.sqlc.GetRoleById(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepo) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) error {
	// TODO: Implement upsert role logic
	return nil
}

func (r *roleRepo) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (bool, error) {
	roleId, err := utils.ToUUID(req.RoleId)
	if err != nil {
		return false, err
	}
	count, err := r.sqlc.DeleteRole(ctx, roleId)

	if err != nil {
		println("Failed to delete role from database:", err)
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (r *roleRepo) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (bool, error) {
	roleId, err := utils.ToUUID(req.RoleId)
	if err != nil {
		return false, err
	}

	count, err := r.sqlc.DisableOrEnableRole(ctx, roleId)
	if err != nil {
		r.logger.ErrorString("failed to disable or enable role in database")
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
