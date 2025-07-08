-- +goose Up
-- +goose StatementBegin
ALTER TABLE roles
ADD COLUMN is_root BOOLEAN NOT NULL DEFAULT FALSE;
AlTER TABLE roles
ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE roles
DROP COLUMN is_root;
ALTER TABLE roles
DROP COLUMN is_active;
-- +goose StatementEnd
