-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ALTER COLUMN lock_end TYPE TIMESTAMPTZ,
    ALTER COLUMN last_login_at TYPE TIMESTAMPTZ,
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE roles
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE resources
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE actions
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE permissions
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE permission_actions
    ALTER COLUMN assign_at TYPE TIMESTAMPTZ;

ALTER TABLE user_roles
    ALTER COLUMN assign_at TYPE TIMESTAMPTZ;

ALTER TABLE role_permissions
    ALTER COLUMN assign_at TYPE TIMESTAMPTZ;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ALTER COLUMN lock_end TYPE TIMESTAMP,
    ALTER COLUMN last_login_at TYPE TIMESTAMP,
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP;

ALTER TABLE roles
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP;

ALTER TABLE resources
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP;

ALTER TABLE actions
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP;

ALTER TABLE permissions
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP;

ALTER TABLE permission_actions
    ALTER COLUMN assign_at TYPE TIMESTAMP;

ALTER TABLE user_roles
    ALTER COLUMN assign_at TYPE TIMESTAMP;

ALTER TABLE role_permissions
    ALTER COLUMN assign_at TYPE TIMESTAMP;
-- +goose StatementEnd
