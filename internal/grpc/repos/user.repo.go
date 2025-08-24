package repos

import (
	"auth_service/internal/grpc/database"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"context"

	"github.com/thanvuc/go-core-lib/log"
)

type userRepo struct {
	sqlc   *database.Queries
	logger log.Logger
}

func (ur *userRepo) GetUsers(ctx context.Context, req *auth.GetUsersRequest) ([]database.GetUsersRow, int32, int32, error) {
    pagination := utils.ToPagination(req.PageQuery)

    users, err := ur.sqlc.GetUsers(
        ctx,
        database.GetUsersParams{
            Column1: req.Search,
            Column2: pagination.Limit,
            Column3: pagination.Offset,
        },
    )
    if err != nil {
        return nil, 0, 0, err
    }

    total, err := ur.sqlc.CountTotalUsers(ctx, req.Search)
    if err != nil {
        return nil, 0, 0, err
    }

    

    return users, int32(total), pagination.Limit, nil
}

