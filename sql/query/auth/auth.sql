-- name: InsertUser :one
INSERT INTO users (email, password_hash, last_login_at)
VALUES ($1, $2, $3)
RETURNING user_id, email, created_at, updated_at;

-- name: InsertExternalProvider :one
INSERT INTO external_provider (sub, provider, user_id)
VALUES ($1, $2, $3)
RETURNING sub;
