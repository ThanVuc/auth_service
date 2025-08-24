-- +goose Up
-- +goose StatementBegin
CREATE TABLE external_provider (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    sub VARCHAR(255) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT external_table_sub_provider_unique UNIQUE (sub, provider)
);
    
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS external_table CASCADE;
-- +goose StatementEnd
