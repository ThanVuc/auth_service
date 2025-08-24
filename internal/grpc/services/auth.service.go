package services

import (
	"auth_service/internal/constant"
	"auth_service/internal/grpc/helper"
	"auth_service/internal/grpc/models"
	"auth_service/internal/grpc/repos"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type authService struct {
	authRepo  repos.AuthRepo
	logger    log.Logger
	jwtHelper helper.JWTHelper
}

func (as *authService) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	resourceIds := make([]string, 0)
	reourceName := make([]string, 0)

	actionIds := make([]string, 0)
	actionName := make([]string, 0)
	actionResourceIds := make([]string, 0)
	items := req.Items

	if len(items) == 0 {
		as.logger.Error("no items to save: error arise at SaveRouteResource/auth.service.go", "", zap.Error(fmt.Errorf("no items to save")))
		return nil, fmt.Errorf("no items to save")
	}

	// get slice
	for _, item := range items {
		if item.Resource.Id != "" {
			resourceIds = append(resourceIds, item.Resource.Id)
			reourceName = append(reourceName, item.Resource.Name)

			for _, action := range item.Actions {
				if action.Id != "" {
					actionIds = append(actionIds, action.Id)
					actionName = append(actionName, action.Name)
					actionResourceIds = append(actionResourceIds, item.Resource.Id)
				}
			}
		}
	}

	// Save resources
	err := as.authRepo.SyncResources(ctx, resourceIds, reourceName)
	if err != nil {
		as.logger.Error("Failed to sync resources", "", zap.Error(err))
		return nil, err
	}

	err = as.authRepo.SyncActions(ctx, actionIds, actionResourceIds, actionName)
	if err != nil {
		as.logger.Error("Failed to sync actions", "", zap.Error(err))
		return nil, err
	}

	resp := &auth.SaveRouteResourceResponse{
		Success: true,
		Message: "Resources and actions saved successfully",
	}

	return resp, nil
}

func (as *authService) LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error) {
	url := "https://www.googleapis.com/oauth2/v3/userinfo"
	googleReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:        utils.RuntimeError(ctx, err),
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	googleReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.GoogleAccessToken))
	client := &http.Client{}
	resp, err := client.Do(googleReq)
	if err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:        utils.RuntimeError(ctx, err),
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	defer resp.Body.Close()
	var userInfo models.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:       utils.RuntimeError(ctx, err),
			AccessToken: "",
		}, err
	}

	if userInfo.Sub == "" {
		return &auth.LoginWithGoogleResponse{
			Error:        utils.RuntimeError(ctx, fmt.Errorf("user info is empty")),
			AccessToken:  "",
			RefreshToken: "",
		}, fmt.Errorf("user info is empty")
	}

	userAccount, roleIDs, err := as.authRepo.LoginWithExternalProvider(ctx, userInfo.Sub, userInfo.Email)
	if err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:        utils.DatabaseError(ctx, err),
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	userID := ""
	roleIds := make([]string, 0)
	// If user exists, use existing user ID and roles; otherwise, register a new user
	if userAccount != nil {
		userID = userAccount.UserID.String()
		for _, roleId := range roleIDs {
			roleIds = append(roleIds, roleId.String())
		}
	} else {
		userIDRegister, roleId, err := as.authRepo.RegisterUserWithExternalProvider(ctx, userInfo, constant.GoogleProvider)
		if err != nil {
			return &auth.LoginWithGoogleResponse{
				Error:        utils.RuntimeError(ctx, err),
				AccessToken:  "",
				RefreshToken: "",
			}, err
		}
		roleIds = append(roleIds, roleId)
		userID = userIDRegister
	}

	// generate jwt token
	accessToken, err := as.jwtHelper.GenerateAccessToken(userID, userInfo.Email, roleIds)
	if err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:       utils.RuntimeError(ctx, err),
			AccessToken: "",
		}, err
	}
	refreshToken := as.jwtHelper.GenerateRefreshToken()
	err = as.jwtHelper.WriteRefreshTokenToRedis(refreshToken)
	if err != nil {
		return &auth.LoginWithGoogleResponse{
			Error:       utils.RuntimeError(ctx, err),
			AccessToken: "",
		}, err
	}

	return &auth.LoginWithGoogleResponse{
		Error:        nil,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *authService) Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error) {
	// TODO: Implement logout logic
	return &common.EmptyResponse{
		Success: utils.ToBoolPointer(true),
		Message: utils.ToStringPointer("Logout successful"),
	}, nil
}
