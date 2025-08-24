-- +goose Up
-- +goose StatementBegin
CREATE TABLE outbox (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    aggregate_type VARCHAR(100) NOT NULL,
    aggregate_id VARCHAR(100) NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    payload JSONB NOT NULL,
    status SMALLINT NOT NULL DEFAULT 1,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    processed_at TIMESTAMPTZ,
    error_message TEXT,
    retry_count INT NOT NULL DEFAULT 0,
    request_id UUID NOT NULL
);
CREATE INDEX idx_outbox_status ON outbox (status);
CREATE INDEX idx_outbox_occurred_at ON outbox (occurred_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_outbox_status;
DROP INDEX IF EXISTS idx_outbox_occurred_at;
DROP TABLE IF EXISTS outbox;
-- +goose StatementEnd
