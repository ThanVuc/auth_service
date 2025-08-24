package services

import (
	"auth_service/internal/grpc/mapper"
	"auth_service/internal/grpc/repos"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thanvuc/go-core-lib/log"
)

type permissionService struct {
	permissionRepo   repos.PermissionRepo
	permissionMapper mapper.PermissionMapper
	logger           log.Logger
	pool             *pgxpool.Pool
}

func (ps *permissionService) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	dbResource, err := ps.permissionRepo.GetResources(ctx)
	if err != nil {
		return &auth.GetResourcesResponse{
			Error:     utils.DatabaseError(ctx, err),
			Resources: nil,
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
			Error:   utils.DatabaseError(ctx, err),
			Actions: nil,
		}, err
	}
	actionItem := ps.permissionMapper.ConvertDbActionsRowToGrpcActions(dbActions)
	resp := &auth.GetActionsResponse{
		Actions: actionItem,
	}

	return resp, nil
}

func (ps *permissionService) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	permissions, totalPerms, totalRoot, err := ps.permissionRepo.GetPermissions(ctx, req)
	if err != nil {
		return &auth.GetPermissionsResponse{
			Error:             utils.DatabaseError(ctx, err),
			Permissions:       nil,
			TotalPermissions:  0,
			Root:              0,
			NonRoot:           0,
			RootPercentage:    0,
			NonRootPercentage: 0,
			PageInfo:          utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalPerms),
		}, fmt.Errorf("failed to get permissions: %w", err)
	}

	if totalPerms == 0 {
		return &auth.GetPermissionsResponse{
			Permissions:       nil,
			TotalPermissions:  0,
			Root:              0,
			NonRoot:           0,
			RootPercentage:    0,
			NonRootPercentage: 0,
			PageInfo:          utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalPerms),
		}, err
	}

	nonRoot := totalPerms - totalRoot
	rootPercentage := utils.RoundToTwoDecimal((float64(totalRoot/totalPerms) * 100))
	nonRootPercentage := 100 - rootPercentage

	permissionsItem := ps.permissionMapper.ConvertDbPermissionsRowToGrpcPermissions(permissions)
	resp := &auth.GetPermissionsResponse{
		Permissions:       permissionsItem,
		TotalPermissions:  totalPerms,
		Root:              totalRoot,
		NonRoot:           nonRoot,
		RootPercentage:    float64(rootPercentage),
		NonRootPercentage: float64(nonRootPercentage),
		PageInfo:          utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalPerms),
	}

	return resp, nil
}

func (ps *permissionService) GetPermission(ctx context.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error) {
	permission, err := ps.permissionRepo.GetPermission(ctx, req)
	if err != nil {
		return &auth.GetPermissionResponse{
			Error:      utils.DatabaseError(ctx, err),
			Permission: nil,
		}, err
	}

	resp := &auth.GetPermissionResponse{
		Permission: ps.permissionMapper.ConvertDbPermissionRowToGrpcPermission(permission),
	}

	return resp, nil
}

func (ps *permissionService) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	if ps.pool == nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("database pool is nil")),
		}, fmt.Errorf("database pool is nil")
	}

	tx, err := ps.pool.Begin(ctx)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("failed to begin transaction: %w", err)),
		}, err
	}
	defer tx.Rollback(ctx)

	permId, err := ps.permissionRepo.UpsertPermission(ctx, tx, req)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("failed to upsert permission: %w", err)),
		}, err
	}

	dbActions, err := ps.permissionRepo.GetActionsByPermissionId(ctx, tx, *permId)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("failed to get actions by permission ID: %w", err)),
		}, err
	}

	addActionIds := utils.Difference(req.ActionIds, dbActions)
	deleteActionIds := utils.Difference(dbActions, req.ActionIds)

	err = ps.permissionRepo.UpdateActionsToPermission(ctx, tx, *permId, addActionIds, deleteActionIds)
	if err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("failed to update actions to permission: %w", err)),
		}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &auth.UpsertPermissionResponse{
			IsSuccess:    false,
			PermissionId: "",
			Error:        utils.DatabaseError(ctx, fmt.Errorf("failed to commit transaction: %w", err)),
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
		return &auth.DeletePermissionResponse{
			Success: false,
			Error:   utils.DatabaseError(ctx, fmt.Errorf("failed to delete permission: %w", err)),
			Message: nil,
		}, fmt.Errorf("failed to delete permission: %w", err)
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
