package mapper

import (
	"auth_service/internal/database"
	"auth_service/proto/auth"
)

type permissionMapper struct{}

func (p *permissionMapper) ConvertDbResourcesRowToGrpcResources(resources []database.GetResourcesRow) []*auth.Resource {
	result := make([]*auth.Resource, 0)
	for _, resource := range resources {
		result = append(result, &auth.Resource{
			Id:   resource.ResourceID,
			Name: resource.Name,
		})
	}

	return result
}

func (p *permissionMapper) ConvertDbActionsRowToGrpcActions(resources []database.GetActionsRow) []*auth.Action {
	result := make([]*auth.Action, 0)
	for _, resource := range resources {
		result = append(result, &auth.Action{
			Id:   resource.ActionID,
			Name: resource.Name,
		})
	}

	return result
}

func (p *permissionMapper) ConvertDbPermissionsRowToGrpcPermissions(permissions []database.GetPermissionsRow) []*auth.PermissionItem {
	result := make([]*auth.PermissionItem, 0)
	for _, permission := range permissions {
		result = append(result, &auth.PermissionItem{
			PermId:   permission.PermID.String(),
			PermName: permission.Name,
			IsRoot:   permission.IsRoot,
		})
	}

	return result
}
