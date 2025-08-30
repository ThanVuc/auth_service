package mapper

import (
	"auth_service/internal/grpc/database"
	"auth_service/proto/auth"

	"github.com/jackc/pgx/v5/pgtype"
)

type (
	AuthMapper interface {
		ConvertFromUserAuthRowToProto(from []database.GetUserAuthInfoRow) []*auth.PermissionAuthItem
	}

	PermissionMapper interface {
		ConvertDbResourcesRowToGrpcResources(resources []database.GetResourcesRow) []*auth.Resource
		ConvertDbActionsRowToGrpcActions(resources []database.GetActionsRow) []*auth.Action
		ConvertDbPermissionsRowToGrpcPermissions(permissions []database.GetPermissionsRow) []*auth.PermissionItem
		ConvertDbPermissionRowToGrpcPermission(permission *[]database.GetPermissionRow) *auth.PermissionItem
	}

	RoleMapper interface {
		ConvertDbRolesRowToGrpcRoles(roles []database.GetRolesRow, usersCount map[pgtype.UUID]int32) []*auth.RoleItem
		ConvertDbRoleByIdRowToGrpcRole(role *[]database.GetRoleByIdRow) *auth.GetRoleResponse
	}

	UserMapper interface {
		ConvertDbUsersRowToGrpcUsers(users []database.GetUsersRow) []*auth.UserItem
		ConvertDbUserRowToGrpcUser(user *[]database.GetUserRow) *auth.UserItem
	}
)

func NewPermissionMapper() PermissionMapper {
	return &permissionMapper{}
}

func NewRoleMapper() RoleMapper {
	return &roleMapper{}
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func NewAuthMapper() AuthMapper {
	return &authMapper{}
}
