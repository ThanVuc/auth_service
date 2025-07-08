package services

import (
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/internal/utils"
	"auth_service/proto/auth"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type roleService struct {
	roleRepo   repos.RoleRepo
	roleMapper mapper.RoleMapper
}

func (rs *roleService) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	roles, totalRoles, total_root, err := rs.roleRepo.GetRoles(ctx, req)

	if err != nil {
		return &auth.GetRolesResponse{
			Error: utils.DatabaseError(err),
		}, err
	}

	if totalRoles == 0 {
		return &auth.GetRolesResponse{
			Roles:      nil,
			TotalRoles: 0,
			NonRoot:    0,
			Error:      nil,
		}, nil
	}

	roleIds := make([]pgtype.UUID, 0, len(roles))
	for _, role := range roles {
		roleIds = append(roleIds, role.RoleID)
	}

	userCounts, err := rs.roleRepo.CountUsersByRoles(ctx, roleIds)
	if err != nil {
		return &auth.GetRolesResponse{
			Roles:      nil,
			TotalRoles: 0,
			NonRoot:    0,
			Error:      utils.DatabaseError(err),
		}, err
	}

	roleCountsMap := make(map[pgtype.UUID]int32, 0)
	for _, userCount := range *userCounts {
		roleCountsMap[userCount.RoleID] = int32(userCount.TotalUsers)
	}

	nonRoot := totalRoles - total_root
	root_percentage := utils.RoundToTwoDecimal((float64(total_root/totalRoles) * 100))
	non_root_percentage := 100 - root_percentage
	resp := &auth.GetRolesResponse{
		Roles:             rs.roleMapper.ConvertDbRolesRowToGrpcRoles(roles, roleCountsMap),
		TotalRoles:        totalRoles,
		NonRoot:           nonRoot,
		Root:              total_root,
		RootPercentage:    root_percentage,
		NonRootPercentage: non_root_percentage,
		PageInfo:          utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalRoles),
	}

	return resp, nil
}

func (rs *roleService) GetRoleById(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error) {
	roles, err := rs.roleRepo.GetRoleById(ctx, req)
	if err != nil {
		return &auth.GetRoleResponse{
			Role:  nil,
			Error: utils.DatabaseError(err),
		}, err
	}

	if len(*roles) == 0 {
		return &auth.GetRoleResponse{
			Role: nil,
		}, nil
	}

	resp := rs.roleMapper.ConvertDbRoleByIdRowToGrpcRole(roles)

	return resp, nil
}

func (rs *roleService) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error) {
	// TODO: Implement upsert role logic
	return &auth.UpsertRoleResponse{}, nil
}

func (rs *roleService) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	isSuccess, err := rs.roleRepo.DeleteRole(ctx, req)
	if err != nil {
		return &auth.DeleteRoleResponse{
			Error: utils.DatabaseError(err),
		}, err
	}

	if !isSuccess {
		msg := "Role not found or cannot be deleted"
		return &auth.DeleteRoleResponse{
			Success: false,
			Message: &msg,
		}, nil
	}

	return &auth.DeleteRoleResponse{
		Success: true,
	}, nil
}

func (rs *roleService) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	isSuccess, err := rs.roleRepo.DisableOrEnableRole(ctx, req)
	println("isSuccess:", isSuccess)

	if err != nil {
		return &auth.DisableOrEnableRoleResponse{
			Error: utils.DatabaseError(err),
		}, err
	}
	if !isSuccess {
		msg := "Role not found or cannot be disabled/enabled"
		return &auth.DisableOrEnableRoleResponse{
			Success: false,
			Message: &msg,
		}, nil
	}

	return &auth.DisableOrEnableRoleResponse{
		Success: true,
		Message: nil,
	}, nil
}

func (rs *roleService) GetRole(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error) {
	return &auth.GetRoleResponse{}, nil
}
