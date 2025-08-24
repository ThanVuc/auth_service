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
		result = append(result, &auth.UserItem{
			UserId:     user.UserID.String(),
			Email:      user.Email,
			LockReason: user.LockReason.String,
			LockEnd:    lockEnd,
		})
	}

	return result
}
