package repos

import (
	"auth_service/internal/database"
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"
)

type authRepo struct {
	sqlc   *database.Queries
	logger *loggers.LoggerZap
}

// All the below methods are for testing purposes only
func (ar *authRepo) SyncResources(ctx context.Context, ids []string, names []string) error {
	err := ar.sqlc.UpsertResources(ctx, database.UpsertResourcesParams{
		Column1: ids,
		Column2: names,
	})

	if err != nil {
		ar.logger.ErrorString("Failed to upsert resource at SyncResources in auth.repo")
		return err
	}

	err = ar.sqlc.RemoveOldResources(ctx, ids)
	if err != nil {
		ar.logger.ErrorString("Failed to delete resources not in use at SyncResources in auth.repo")
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
		ar.logger.ErrorString("Failed to upsert action at SyncActions in auth.repo")
		return err
	}

	err = ar.sqlc.RemoveOldActions(ctx, ids)
	if err != nil {
		ar.logger.ErrorString("Failed to delete actions not in use at SyncActions in auth.repo")
		return err
	}

	return nil
}

func (ar *authRepo) Login(ctx context.Context, req *auth.LoginRequest) error {
	// TODO: Implement login repository logic
	return nil
}

func (ar *authRepo) Register(ctx context.Context, req *auth.RegisterRequest) error {
	// TODO: Implement register repository logic
	return nil
}

func (ar *authRepo) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) error {
	// TODO: Implement confirm email repository logic
	return nil
}

func (ar *authRepo) Logout(ctx context.Context, req *auth.LogoutRequest) error {
	// TODO: Implement logout repository logic
	return nil
}

func (ar *authRepo) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) error {
	// TODO: Implement forgot password repository logic
	return nil
}

func (ar *authRepo) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) error {
	// TODO: Implement confirm forgot password repository logic
	return nil
}

func (ar *authRepo) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) error {
	// TODO: Implement reset password repository logic
	return nil
}
