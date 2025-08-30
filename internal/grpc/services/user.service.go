package services

import (
	"auth_service/internal/grpc/mapper"
	"auth_service/internal/grpc/repos"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
	"fmt"

	"github.com/thanvuc/go-core-lib/log"
)

type userService struct {
	userRepo   repos.UserRepo
	userMapper mapper.UserMapper
	logger     log.Logger
}

func (ps *userService) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	users, totalUsers, limit, err := ps.userRepo.GetUsers(ctx, req)
	if err != nil {
		return &auth.GetUsersResponse{
			Error:      utils.DatabaseError(ctx, err),
			Users:      nil,
			TotalUsers: 0,
			PageInfo:   utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalUsers),
		}, fmt.Errorf("failed to get users: %w", err)
	}

	if totalUsers == 0 {
		return &auth.GetUsersResponse{
			Error:      nil,
			Users:      nil,
			TotalUsers: 0,
			PageInfo:   utils.ToPageInfo(req.PageQuery.Page, req.PageQuery.PageSize, totalUsers),
		}, err
	}

	usersIterm := ps.userMapper.ConvertDbUsersRowToGrpcUsers(users)
	resp := &auth.GetUsersResponse{
		Users:      usersIterm,
		TotalUsers: totalUsers,
		PageInfo:   utils.ToPageInfo(req.PageQuery.Page, limit, totalUsers),
	}

	return resp, nil
}

func (us *userService) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleToUserRequest) (*common.EmptyResponse, error) {
	err := us.userRepo.AssignRoleToUser(ctx, req)
	if err != nil {
		return &common.EmptyResponse{
			Success: utils.ToBoolPointer(false),
			Message: utils.ToStringPointer("Failed to assign role to user"),
			Error:   utils.DatabaseError(ctx, err),
		}, err
	}

	return &common.EmptyResponse{
		Success: utils.ToBoolPointer(true),
		Message: utils.ToStringPointer("Role assigned successfully"),
	}, nil
}

func (ps *userService) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	user, err := ps.userRepo.GetUser(ctx, req)
	if err != nil {
		return &auth.GetUserResponse{
			Error: utils.DatabaseError(ctx, err),
			User:  nil,
		}, err
	}

	resp := &auth.GetUserResponse{
		User: ps.userMapper.ConvertDbUserRowToGrpcUser(user),
	}
	return resp, nil
}

func (us *userService) LockOrUnLockUser(ctx context.Context, req *auth.LockUserRequest) (*common.EmptyResponse, error) {
    locked, err := us.userRepo.LockOrUnLockUser(ctx, req)
    if err != nil {
        return &common.EmptyResponse{
            Success: utils.ToBoolPointer(false),
            Message: utils.ToStringPointer("Failed to lock/unlock user!"),
            Error:   utils.DatabaseError(ctx, err),
        }, err
    }

    var msg string
    if locked {
        msg = "User locked successfully!"
    } else {
        msg = "User unlocked successfully!"
    }

    return &common.EmptyResponse{
        Success: utils.ToBoolPointer(true),
        Message: utils.ToStringPointer(msg),
    }, nil
}
