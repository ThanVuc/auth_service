-- +goose Up
-- +goose StatementBegin
ALTER TABLE permissions
ADD COLUMN is_root BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE permissions
DROP COLUMN is_root;
-- +goose StatementEnd
