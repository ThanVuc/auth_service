-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id UUID PRIMARY KEY NOT NULL,   -- Tự gán ID
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
    role_id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE permissions (
    perm_id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    action VARCHAR(50) NOT NULL,
    resource VARCHAR(100) NOT NULL,
    description TEXT,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_roles (
    ur_id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    assign_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id) ON DELETE CASCADE
);

CREATE TABLE role_permissions (
    rp_id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    role_id UUID NOT NULL,
    perm_id UUID NOT NULL,
    assign_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(role_id) ON DELETE CASCADE,
    CONSTRAINT fk_perm FOREIGN KEY(perm_id) REFERENCES permissions(perm_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
