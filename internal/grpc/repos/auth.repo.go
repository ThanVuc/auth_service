package repos

import (
	"auth_service/internal/constant"
	"auth_service/internal/grpc/database"
	"auth_service/internal/grpc/models"
	"auth_service/internal/grpc/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type authRepo struct {
	sqlc   *database.Queries
	logger log.Logger
	pool   *pgxpool.Pool
}

// All the below methods are for testing purposes only
func (ar *authRepo) SyncResources(ctx context.Context, ids []string, names []string) error {
	err := ar.sqlc.UpsertResources(ctx, database.UpsertResourcesParams{
		Column1: ids,
		Column2: names,
	})

	if err != nil {
		return err
	}

	err = ar.sqlc.RemoveOldResources(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepo) SyncActions(ctx context.Context, ids, resourceIds, names []string) error {
	err := ar.sqlc.UpsertActions(ctx, database.UpsertActionsParams{
		Column1: ids,
		Column2: resourceIds,
		Column3: names,
	})

	if err != nil {
		return err
	}

	err = ar.sqlc.RemoveOldActions(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepo) RegisterUserWithExternalProvider(ctx context.Context, userInfo models.GoogleUserInfo, provider constant.Provider) (string, string, error) {
	requestID := utils.GetRequestIDFromOutgoingContext(ctx)
	tx, err := ar.pool.Begin(ctx)
	if err != nil {
		ar.logger.Error("Failed to begin transaction", requestID, zap.Error(err))
		return "", "", err
	}
	defer tx.Rollback(ctx)
	qtx := ar.sqlc.WithTx(tx)

	// Insert user into the database
	rowResp, err := qtx.InsertUser(ctx, database.InsertUserParams{
		Email:        userInfo.Email,
		PasswordHash: "",
		LastLoginAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})

	if err != nil || !rowResp.UserID.Valid {
		ar.logger.Error("Failed to insert user", requestID, zap.Error(err))
		tx.Rollback(ctx)
		return "", "", err
	}

	// Insert external provider
	_, err = qtx.InsertExternalProvider(ctx, database.InsertExternalProviderParams{
		Sub:      userInfo.Sub,
		Provider: string(provider),
		UserID:   rowResp.UserID,
	})

	if err != nil {
		ar.logger.Error("Failed to insert external provider", requestID, zap.Error(err))
		tx.Rollback(ctx)
		return "", "", err
	}

	// Assign Default role to the new user
	defaultRole, err := qtx.GetRoleByName(ctx, string(constant.UserRole))
	if err != nil {
		ar.logger.Error("Failed to get default role", requestID, zap.Error(err))
		tx.Rollback(ctx)
		return "", "", err
	}

	if !defaultRole.RoleID.Valid || defaultRole.RoleID.String() == "" {
		tx.Rollback(ctx)
		return "", "", fmt.Errorf("default role not found")
	}

	err = qtx.AddNewRolesToUser(ctx, database.AddNewRolesToUserParams{
		UserID:  rowResp.UserID,
		Column2: []pgtype.UUID{defaultRole.RoleID},
	})

	if err != nil {
		ar.logger.Error("Failed to assign default role to user", requestID, zap.Error(err))
		tx.Rollback(ctx)
		return "", "", err
	}

	// Insert to outbox to sync with other services
	var outboxPayload = map[string]interface{}{
		"user_id":    rowResp.UserID.String(),
		"email":      rowResp.Email,
		"created_at": rowResp.CreatedAt.Time.Unix(),
		"name":       userInfo.Name,
	}

	requestId := utils.GetRequestIDFromOutgoingContext(ctx)
	payloadBytes, marshalErr := json.Marshal(outboxPayload)
	if marshalErr != nil {
		ar.logger.Error("Failed to marshal outbox payload for new user", requestID, zap.Error(marshalErr))
		tx.Rollback(ctx)
		return "", "", marshalErr
	}
	_, err = qtx.InsertOutbox(ctx, database.InsertOutboxParams{
		AggregateType: constant.AggregateTypeUser,
		AggregateID:   rowResp.UserID.String(),
		EventType:     constant.EventTypeCreate,
		Payload:       payloadBytes,
		RequestID:     requestId,
	})
	if err != nil {
		ar.logger.Error("Failed to insert outbox for new user", requestID, zap.Error(err))
		tx.Rollback(ctx)
		return "", "", err
	}

	if err = tx.Commit(ctx); err != nil {
		ar.logger.Error("Failed to commit transaction", "", zap.Error(err))
		return "", "", err
	}

	return rowResp.UserID.String(), defaultRole.RoleID.String(), nil
}

func (ar *authRepo) LoginWithExternalProvider(ctx context.Context, sub string, email string) (*database.LoginWithExternalProviderRow, []pgtype.UUID, error) {
	userAccount, err := ar.sqlc.LoginWithExternalProvider(ctx, database.LoginWithExternalProviderParams{
		Sub:   sub,
		Email: email,
	})

	if err != nil && err != pgx.ErrNoRows {
		return nil, nil, err
	}

	if !userAccount.UserID.Valid || userAccount.UserID.String() == "" || err == pgx.ErrNoRows {
		return nil, nil, nil
	}

	err = ar.sqlc.UpdateUserLastLogin(ctx, userAccount.UserID)

	if err != nil {
		return nil, nil, err
	}

	userRoleIDs, err := ar.sqlc.GetRoleIDsByUserID(ctx, userAccount.UserID)
	if err != nil {
		return nil, nil, err
	}

	if len(userRoleIDs) == 0 {
		return nil, nil, fmt.Errorf("user has no roles assigned")
	}

	return &userAccount, userRoleIDs, nil
}

func (ar *authRepo) CheckPermission(ctx context.Context, roleIDs []string, resource string, action string) (bool, error) {
	// Convert roleIDs from []string to []pgtype.UUID
	pgRoleIDs := make([]pgtype.UUID, len(roleIDs))
	for i, id := range roleIDs {
		var err error
		pgRoleIDs[i], err = utils.ToUUID(id)
		if err != nil {
			return false, fmt.Errorf("invalid role ID: %s", id)
		}
	}

	hasPermission, err := ar.sqlc.HasPermission(ctx, database.HasPermissionParams{
		Column1: pgRoleIDs,
		Name:    resource,
		Name_2:  action,
	})

	if err != nil {
		return false, err
	}

	return hasPermission, nil
}

func (ar *authRepo) GetUserActionsAndResources(ctx context.Context, roleIDs []string) ([]database.GetUserAuthInfoRow, error) {
	// Convert roleIDs from []string to []pgtype.UUID
	pgRoleIDs := make([]pgtype.UUID, len(roleIDs))
	for i, id := range roleIDs {
		var err error
		pgRoleIDs[i], err = utils.ToUUID(id)
		if err != nil {
			return nil, fmt.Errorf("invalid role ID: %s", id)
		}
	}

	rows, err := ar.sqlc.GetUserAuthInfo(ctx, pgRoleIDs)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (ar *authRepo) SyncDatabase(ctx context.Context) error {
	users, err := ar.sqlc.GetUsers(ctx, database.GetUsersParams{})
	if err != nil {
		return err
	}

	requestIds := make([]string, 0)
	payloads := make([][]byte, 0)
	aggregate_types := make([]string, 0)
	aggregate_ids := make([]string, 0)
	event_types := make([]string, 0)
	requestId := utils.GetRequestIDFromOutgoingContext(ctx)

	for _, user := range users {
		var outboxPayload = map[string]interface{}{
			"user_id":    user.UserID.String(),
			"email":      user.Email,
			"created_at": time.Now().Unix(),
		}

		payloadBytes, marshalErr := json.Marshal(outboxPayload)
		if marshalErr != nil {
			return marshalErr
		}
		requestIds = append(requestIds, requestId)
		payloads = append(payloads, payloadBytes)
		aggregate_types = append(aggregate_types, constant.AggregateTypeUser)
		aggregate_ids = append(aggregate_ids, user.UserID.String())
		event_types = append(event_types, constant.EventTypeCreate)
	}

	err = ar.sqlc.InsertOutboxBulk(ctx, database.InsertOutboxBulkParams{
		Column1: aggregate_types,
		Column2: aggregate_ids,
		Column3: event_types,
		Column4: payloads,
		Column5: requestIds,
	})

	if err != nil {
		return err
	}

	return nil
}
