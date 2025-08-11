package repos

import (
	"auth_service/proto/auth"
	"context"

	"github.com/thanvuc/go-core-lib/log"
)

type tokenRepo struct {
	logger log.Logger
}

func (r *tokenRepo) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) error {
	// TODO: Implement refresh token logic
	return nil
}

func (r *tokenRepo) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) error {
	// TODO: Implement revoke token logic
	return nil
}
