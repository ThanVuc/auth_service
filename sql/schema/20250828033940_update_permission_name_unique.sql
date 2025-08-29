-- +goose Up
-- +goose StatementBegin
Alter table permission
    Drop constraint if exists permission_name_unique,
    Add constraint permission_name_unique unique (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter table permission
    Drop constraint if exists permission_name_unique;
-- +goose StatementEnd
