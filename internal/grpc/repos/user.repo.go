package repos

import (
	"auth_service/internal/grpc/database"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"context"
	"time"

	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thanvuc/go-core-lib/log"
)

type userRepo struct {
	sqlc   *database.Queries
	logger log.Logger
	pool   *pgxpool.Pool
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

func (r *userRepo) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleToUserRequest) error {
	userID, err := utils.ToUUID(req.UserId)
	if err != nil {
		return err
	}
	oldRoleIDs, err := r.GetRoleIDsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	newRoleIDs := make([]pgtype.UUID, 0, len(req.RoleIds))
	for _, roleId := range req.RoleIds {
		uuid, err := utils.ToUUID(roleId)
		if err != nil {
			return err
		}
		newRoleIDs = append(newRoleIDs, uuid)
	}

	addRoleIDs := utils.Difference(newRoleIDs, oldRoleIDs)
	removeRoleIDs := utils.Difference(oldRoleIDs, newRoleIDs)
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	if len(addRoleIDs) != 0 {
		err = r.AddNewRolesToUser(ctx, tx, userID, addRoleIDs)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	if len(removeRoleIDs) != 0 {
		err = r.RemoveRolesFromUser(ctx, tx, userID, removeRoleIDs)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetRoleIDsByUserID(ctx context.Context, userId pgtype.UUID) ([]pgtype.UUID, error) {
	roleIDs, err := r.sqlc.GetRoleIDsByUserID(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		requestID := utils.GetRequestIDFromOutgoingContext(ctx)
		r.logger.Error("Failed to get role IDs by user ID", requestID)
		return nil, err
	}
	return roleIDs, nil
}

func (r *userRepo) AddNewRolesToUser(ctx context.Context, tx pgx.Tx, userId pgtype.UUID, ids []pgtype.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	err := r.sqlc.AddNewRolesToUser(ctx, database.AddNewRolesToUserParams{
		UserID:  userId,
		Column2: ids,
	})

	if err != nil {
		requestID := utils.GetRequestIDFromOutgoingContext(ctx)
		r.logger.Error("Failed to add new roles to user", requestID)
		return err
	}

	return nil
}

func (r *userRepo) RemoveRolesFromUser(ctx context.Context, tx pgx.Tx, userId pgtype.UUID, ids []pgtype.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	err := r.sqlc.RemoveRolesFromUser(ctx, database.RemoveRolesFromUserParams{
		UserID:  userId,
		Column2: ids,
	})

	if err != nil {
		requestID := utils.GetRequestIDFromOutgoingContext(ctx)
		r.logger.Error("Failed to remove roles from user", requestID)
		return err
	}

	return nil
}
func (ur *userRepo) GetUser(ctx context.Context, req *auth.GetUserRequest) (*[]database.GetUserRow, error) {
	userIdUUID, err := utils.ToUUID(req.UserId)
	if err != nil {
		return nil, err
	}

	user, err := ur.sqlc.GetUser(ctx, userIdUUID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &user, nil

}

func (ur *userRepo) LockOrUnLockUser(ctx context.Context, req *auth.LockUserRequest) error {
	userID, err := utils.ToUUID(req.UserId)
	if err != nil {
		return err
	}

	var lockEnd pgtype.Timestamptz
	lockEnd, err = ur.sqlc.GetLockEndByUserID(ctx, userID)
	if err != nil {
		return err
	}
	
	now := time.Now()
	if lockEnd.Valid && lockEnd.Time.After(now) {
		return ur.sqlc.UnlockUser(ctx, userID)
	}

	err = ur.sqlc.LockUser(ctx, database.LockUserParams{
		UserID:     userID,
		LockReason: pgtype.Text{String: *req.LockReason, Valid: true},
	})
	if err != nil {
		return err

	}

	return nil

}
