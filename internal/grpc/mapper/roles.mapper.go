package mapper

import (
	"auth_service/internal/grpc/database"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"

	"github.com/jackc/pgx/v5/pgtype"
)

type roleMapper struct{}

func (r *roleMapper) ConvertDbRolesRowToGrpcRoles(roles []database.GetRolesRow, usersCount map[pgtype.UUID]int32) []*auth.RoleItem {
	result := make([]*auth.RoleItem, 0, len(roles))
	for _, role := range roles {
		result = append(result, &auth.RoleItem{
			RoleId:      role.RoleID.String(),
			Name:        role.Name,
			Description: role.Description.String,
			IsRoot:      role.IsRoot,
			TotalUsers:  usersCount[role.RoleID],
			IsActive:    role.IsActive,
		})
	}
	return result
}

func (r *roleMapper) ConvertDbRoleByIdRowToGrpcRole(role *[]database.GetRoleByIdRow) *auth.GetRoleResponse {
	if role == nil || len(*role) == 0 {
		return nil
	}

	perms := make([]*auth.PermissionItem, 0, len((*role)))
	for _, r := range *role {
		perms = append(perms, &auth.PermissionItem{
			PermId:      r.PermissionID.String(),
			PermName:    r.RoleName,
			Description: r.Description.String,
		})
	}

	return &auth.GetRoleResponse{
		Role: &auth.RoleItem{
			RoleId:      (*role)[0].RoleID.String(),
			Name:        (*role)[0].RoleName,
			Description: (*role)[0].Description.String,
			IsRoot:      (*role)[0].IsRoot,
			CreatedAt:   utils.FromPgTypeTimeStamptZToUnix((*role)[0].CreatedAt),
			UpdatedAt:   utils.FromPgTypeTimeStamptZToUnix((*role)[0].UpdatedAt),
			IsActive:    (*role)[0].IsActive,
			Permissions: perms,
		},
	}
}
