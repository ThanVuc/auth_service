-- +goose Up
-- +goose StatementBegin
ALTER TABLE outbox
ALTER COLUMN request_id TYPE VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE outbox
ALTER COLUMN request_id TYPE UUID USING request_id::UUID;
-- +goose StatementEnd
