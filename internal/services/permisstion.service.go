package services

import (
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/internal/utils"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgxpool"
)

type permissionService struct {
	permissionRepo   repos.PermissionRepo
	permissionMapper mapper.PermissionMapper
	logger           loggers.LoggerZap
	pool             *pgxpool.Pool
}

func (ps *permissionService) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	dbResource, err := ps.permissionRepo.GetResources(ctx)
	if err != nil {
		return &auth.GetResourcesResponse{
			Error: utils.DatabaseError("Failed to get resources"),
		}, err
	}
	resourceItem := ps.permissionMapper.ConvertDbResourcesRowToGrpcResources(dbResource)
	resp := &auth.GetResourcesResponse{
		Resources: resourceItem,
	}

	return resp, nil
}

func (ps *permissionService) GetActions(ctx context.Context, items *auth.GetActionsRequest) (*auth.GetActionsResponse, error) {
	dbActions, err := ps.permissionRepo.GetActions(ctx, items.ResourceId)
	if err != nil {
		return &auth.GetActionsResponse{
			Error: utils.DatabaseError("Failed to get actions"),
		}, err
	}
	actionItem := ps.permissionMapper.ConvertDbActionsRowToGrpcActions(dbActions)
	resp := &auth.GetActionsResponse{
		Actions: actionItem,
	}

	return resp, nil
}

func (ps *permissionService) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	permissions, total_perms, total_root, err := ps.permissionRepo.GetPermissions(ctx, req)
	if err != nil {
		return &auth.GetPermissionsResponse{
			Error: utils.DatabaseError("Failed to get permissions"),
		}, nil
	}

	if total_perms == 0 {
		return &auth.GetPermissionsResponse{
			Permissions:       nil,
			TotalPermissions:  0,
			Root:              0,
			NonRoot:           0,
			RootPercentage:    0,
			NonRootPercentage: 0,
		}, nil
	}

	nonRoot := total_perms - total_root
	root_percentage := ps.roundToTwoDecimal((float64(total_root/total_perms) * 100))
	non_root_percentage := 100 - root_percentage

	permissionsItem := ps.permissionMapper.ConvertDbPermissionsRowToGrpcPermissions(permissions)
	resp := &auth.GetPermissionsResponse{
		Permissions:       permissionsItem,
		TotalPermissions:  total_perms,
		Root:              total_root,
		NonRoot:           nonRoot,
		RootPercentage:    float64(root_percentage),
		NonRootPercentage: float64(non_root_percentage),
		PageInfo:          utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, total_perms),
	}

	return resp, nil
}

func (ps *permissionService) roundToTwoDecimal(val float64) float64 {
	return math.Round(val*100) / 100
}

func (ps *permissionService) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	tx, err := ps.pool.Begin(ctx)
	if err != nil {
		ps.logger.ErrorString("failed to begin transaction for upserting permission")
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError("Failed to begin transaction"),
		}, err
	}
	defer tx.Rollback(ctx)

	permId, err := ps.permissionRepo.UpsertPermission(ctx, tx, req)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError("Failed to upsert permission"),
		}, err
	}

	dbActions, err := ps.permissionRepo.GetActionsByPermissionId(ctx, tx, *permId)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError("Failed to get actions by permission ID"),
		}, err
	}

	addActionIds := make([]string, 0)
	deleteActionIds := make([]string, 0)
	for _, action := range req.ActionsIds {
		if !utils.Contains(dbActions, action) {
			addActionIds = append(addActionIds, action)
		}
	}

	for _, dbAction := range dbActions {
		if !utils.Contains(req.ActionsIds, dbAction) {
			deleteActionIds = append(deleteActionIds, dbAction)
		}
	}

	err = ps.permissionRepo.UpdateActionsToPermission(ctx, tx, *permId, addActionIds, deleteActionIds)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError("Failed to update actions to permission"),
		}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError("Failed to commit transaction"),
		}, err
	}

	return &auth.UpsertPermissionResponse{
		IsSuccess:    true,
		PermissionId: permId.String(),
		Error:        nil,
	}, nil
}

func (ps *permissionService) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	// TODO: Implement delete permission logic
	return &auth.DeletePermissionResponse{}, nil
}

func (ps *permissionService) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	// TODO: Implement assign permission to role logic
	return &auth.AssignPermissionResponse{}, nil
}
