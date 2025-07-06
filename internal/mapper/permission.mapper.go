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

func (p *permissionMapper) ConvertDbPermissionRowToGrpcPermission(permission *[]database.GetPermissionRow) *auth.PermissionItem {
	resp := &auth.PermissionItem{}
	if permission == nil || len(*permission) == 0 {
		return resp
	}

	actions := make([]*auth.Action, 0, len((*permission)))
	for _, perm := range *permission {
		if perm.ActionID.Valid && perm.ActionName.Valid {
			actions = append(actions, &auth.Action{
				Id:   perm.ActionID.String,
				Name: perm.ActionName.String,
			})
		}
	}

	perm := (*permission)[0]
	resp.PermId = perm.PermID.String()
	resp.PermName = perm.PermissionName
	resp.Resource = &auth.Resource{
		Id:   perm.ResourceID,
		Name: perm.ResourceName,
	}
	resp.Description = perm.Description.String
	resp.IsRoot = perm.IsRoot
	resp.Actions = actions

	if perm.CreatedAt.Valid {
		timestamp := perm.CreatedAt.Time.Unix()
		resp.CreatedAt = &timestamp
	}

	if perm.UpdatedAt.Valid {
		timestamp := perm.UpdatedAt.Time.Unix()
		resp.UpdatedAt = &timestamp
	}

	return resp
}
