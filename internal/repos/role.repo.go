package repos

import (
	"auth_service/internal/database"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
)

type roleRepo struct {
	logger *loggers.LoggerZap
	sqlc   *database.Queries
}

func (r *roleRepo) GetRoles(ctx context.Context, req *auth.GetRolesRequest) error {
	// TODO: Implement get roles logic
	return nil
}

func (r *roleRepo) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) error {
	// TODO: Implement upsert role logic
	return nil
}

func (r *roleRepo) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) error {
	// TODO: Implement delete role logic
	return nil
}

func (r *roleRepo) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) error {
	// TODO: Implement disable/enable role logic
	return nil
}
