package repos

import (
	"auth_service/internal/database"
	"auth_service/internal/utils"
	"auth_service/proto/auth"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thanvuc/go-core-lib/log"
)

type roleRepo struct {
	logger log.Logger
	sqlc   *database.Queries
}

func (r *roleRepo) GetRoles(ctx context.Context, req *auth.GetRolesRequest) ([]database.GetRolesRow, int32, int32, error) {
	pagination := utils.ToPagination(req.PageQuery)

	roles, err := r.sqlc.GetRoles(ctx, database.GetRolesParams{
		Column1: req.Search,
		Column2: pagination.Limit,
		Column3: pagination.Offset,
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

func (r *roleRepo) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (bool, error) {
	roleId, err := utils.ToUUID(req.RoleId)
	if err != nil {
		return false, err
	}
	count, err := r.sqlc.DeleteRole(ctx, roleId)

	if err != nil {
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
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (r *roleRepo) UpsertRole(ctx context.Context, tx pgx.Tx, req *auth.UpsertRoleRequest) (string, error) {
	sqlcTx := r.sqlc.WithTx(tx)
	if req.RoleId == nil {
		roleId, err := sqlcTx.InsertRole(ctx, database.InsertRoleParams{
			Name:        req.Name,
			Description: pgtype.Text{String: req.Description, Valid: true},
		})

		if err != nil {
			return "", err
		}
		return roleId.String(), nil
	}
	roleId, err := utils.ToUUID(*req.RoleId)
	if err != nil {
		return "", err
	}
	rowCount, err := sqlcTx.UpdateRole(ctx, database.UpdateRoleParams{
		RoleID:      roleId,
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: true},
	})

	if err != nil {
		return "", err
	}

	if rowCount == 0 {
		return "", pgx.ErrNoRows
	}

	return roleId.String(), nil
}

func (r *roleRepo) GetPermissionIdsByRole(ctx context.Context, tx pgx.Tx, roleId string) ([]pgtype.UUID, error) {
	roleIdUUID, err := utils.ToUUID(roleId)
	if err != nil {
		return nil, err
	}

	if roleIdUUID == (pgtype.UUID{}) {
		return nil, pgx.ErrNoRows
	}

	permIds, err := r.sqlc.GetPermissionIdsByRole(ctx, roleIdUUID)
	if err != nil {
		return nil, err
	}

	return permIds, nil
}

func (r *roleRepo) UpsertPermissionsForRole(ctx context.Context, tx pgx.Tx, roleId string, addPerms *[]pgtype.UUID, delPerms *[]pgtype.UUID) (bool, error) {
	roleIdUUID, err := utils.ToUUID(roleId)
	if err != nil {
		return false, err
	}

	if roleIdUUID == (pgtype.UUID{}) {
		return false, pgx.ErrNoRows
	}

	sqlcTx := r.sqlc.WithTx(tx)
	if addPerms != nil && len(*addPerms) > 0 {
		count, err := sqlcTx.AddPermissionsToRole(ctx, database.AddPermissionsToRoleParams{
			RoleID:  roleIdUUID,
			Column2: *addPerms,
		})

		if err != nil {
			return false, err
		}

		if count == 0 {
			return false, pgx.ErrNoRows
		}
	}

	if delPerms != nil && len(*delPerms) > 0 {
		count, err := sqlcTx.RemovePermissionsFromRole(ctx, database.RemovePermissionsFromRoleParams{
			RoleID:  roleIdUUID,
			Column2: *delPerms,
		})

		if err != nil {
			return false, err
		}

		if count == 0 {
			return false, pgx.ErrNoRows
		}
	}

	return true, nil

}
