package services

import (
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/internal/utils"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
	"fmt"

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
			Error: utils.DatabaseError(err),
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
			Error: utils.DatabaseError(err),
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
			Error: utils.DatabaseError(err),
		}, err
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
	root_percentage := utils.RoundToTwoDecimal((float64(total_root/total_perms) * 100))
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

func (ps *permissionService) GetPermission(ctx context.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error) {
	permission, err := ps.permissionRepo.GetPermission(ctx, req)
	if err != nil {
		ps.logger.ErrorString("failed to get permission from database")
		return &auth.GetPermissionResponse{
			Error: utils.DatabaseError(err),
		}, err
	}

	resp := &auth.GetPermissionResponse{
		Permission: ps.permissionMapper.ConvertDbPermissionRowToGrpcPermission(permission),
	}

	return resp, nil
}

func (ps *permissionService) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	if ps.pool == nil {
		ps.logger.ErrorString("database pool is nil")
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(fmt.Errorf("database connection not available")),
		}, fmt.Errorf("database pool is nil")
	}

	tx, err := ps.pool.Begin(ctx)
	if err != nil {
		ps.logger.ErrorString("failed to begin transaction for upserting permission")
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(err),
		}, err
	}
	defer tx.Rollback(ctx)

	permId, err := ps.permissionRepo.UpsertPermission(ctx, tx, req)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(err),
		}, err
	}

	dbActions, err := ps.permissionRepo.GetActionsByPermissionId(ctx, tx, *permId)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(err),
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
			Error:        utils.DatabaseError(err),
		}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(err),
		}, err
	}

	return &auth.UpsertPermissionResponse{
		IsSuccess:    true,
		PermissionId: permId.String(),
		Error:        nil,
	}, nil
}

func (ps *permissionService) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	isScuess, err := ps.permissionRepo.DeletePermission(ctx, req)
	if err != nil {
		ps.logger.ErrorString("failed to delete permission from database")
		return &auth.DeletePermissionResponse{
			Success: false,
			Error:   utils.DatabaseError(err),
		}, err
	}

	if !isScuess {
		msg := "Permission not found or cannot be deleted"
		return &auth.DeletePermissionResponse{
			Success: false,
			Message: &msg,
			Error:   nil,
		}, nil
	}

	return &auth.DeletePermissionResponse{
		Success: true,
		Error:   nil,
	}, nil
}
