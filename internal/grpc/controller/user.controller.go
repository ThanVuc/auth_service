package controller

import (
	"auth_service/internal/grpc/services"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
)

type UserController struct {
	auth.UnimplementedUserServiceServer
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (pc *UserController) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.userService.GetUsers)
}

func (uc *UserController) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleToUserRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, req, uc.userService.AssignRoleToUser)
}
