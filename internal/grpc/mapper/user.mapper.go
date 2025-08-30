package mapper

import (
	"auth_service/internal/grpc/database"
	"auth_service/proto/auth"
)

type userMapper struct{}

func (u *userMapper) ConvertDbUsersRowToGrpcUsers(users []database.GetUsersRow) []*auth.UserItem {
	result := make([]*auth.UserItem, 0)
	for _, user := range users {

		var lockEnd int64
		if user.LockEnd.Valid {
			lockEnd = user.LockEnd.Time.Unix()
		} else {
			lockEnd = 0
		}

		var lastLoginAt int64
		if user.LastLoginAt.Valid {
			lastLoginAt = user.LastLoginAt.Time.Unix()
		} else {
			lastLoginAt = 0
		}

		result = append(result, &auth.UserItem{
			UserId:     user.UserID.String(),
			Email:      user.Email,
			LockReason: user.LockReason.String,
			LockEnd:    lockEnd,
			LastLoginAt : lastLoginAt,
		})
	}

	return result
}

func (u *userMapper) ConvertDbUserRowToGrpcUser(user *[]database.GetUserRow) *auth.UserItem {
	resp := &auth.UserItem{}
	if user == nil || len(*user) == 0 {
		return resp
	}

	roles := make([]*auth.RoleItem, 0, len((*user)))
	for _, r := range *user {
		if r.RoleID.Valid && r.RoleName.Valid {
			roles = append(roles, &auth.RoleItem{
				RoleId: r.RoleID.String(),
				Name:   r.RoleName.String,
				Description: r.RoleDescription.String,
			})
		}
	}

	userData := (*user)[0]
	resp.UserId = userData.UserID.String()
	resp.Email = userData.Email
	resp.LockReason = userData.LockReason.String
	resp.Roles = roles
	if userData.LockEnd.Valid {
		resp.LockEnd = userData.LockEnd.Time.Unix()
	} else {
		resp.LockEnd = 0
	}

	if userData.CreatedAt.Valid {
		timestamp := userData.CreatedAt.Time.Unix()
		resp.CreatedAt = timestamp
	}

	if userData.UpdatedAt.Valid {
		timestamp := userData.UpdatedAt.Time.Unix()
		resp.UpdatedAt = timestamp
	}

	if userData.LastLoginAt.Valid {
		timestamp := userData.LastLoginAt.Time.Unix()
		resp.LastLoginAt = timestamp
	}

	return resp

}
