package mapper

import (
	"auth_service/internal/database"
	"auth_service/proto/auth"
)

type (
	PermissionMapper interface {
		ConvertDbResourcesRowToGrpcResources(resources []database.GetResourcesRow) []*auth.Resource
		ConvertDbActionsRowToGrpcActions(resources []database.GetActionsRow) []*auth.Action
		ConvertDbPermissionsRowToGrpcPermissions(permissions []database.GetPermissionsRow) []*auth.PermissionItem
	}
)

func NewPermissionMapper() PermissionMapper {
	return &permissionMapper{}
}
