-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    -- For updated_at field, update on row UPDATE
    IF TG_ARGV[0] = 'updated_at' THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END IF;

    -- For assign_at field, update on INSERT and UPDATE
    IF TG_ARGV[0] = 'assign_at' THEN
        IF TG_OP = 'INSERT' THEN
            NEW.assign_at = CURRENT_TIMESTAMP;
        ELSIF TG_OP = 'UPDATE' THEN
            NEW.assign_at = CURRENT_TIMESTAMP;
        END IF;
        RETURN NEW;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Triggers for updating timestamps
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

CREATE TRIGGER trg_roles_updated_at
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

CREATE TRIGGER trg_resources_updated_at
BEFORE UPDATE ON resources
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

CREATE TRIGGER trg_actions_updated_at
BEFORE UPDATE ON actions
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

CREATE TRIGGER trg_permissions_updated_at
BEFORE UPDATE ON permissions
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

-- Assign_at triggers for user_roles and role_permissions
CREATE TRIGGER trg_permission_actions_assign_at
BEFORE INSERT OR UPDATE ON permission_actions
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('assign_at');

CREATE TRIGGER trg_user_roles_assign_at
BEFORE INSERT OR UPDATE ON user_roles
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('assign_at');

CREATE TRIGGER trg_role_permissions_assign_at
BEFORE INSERT OR UPDATE ON role_permissions
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('assign_at');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop assign_at triggers
DROP TRIGGER IF EXISTS trg_permission_actions_assign_at ON permission_actions;
DROP TRIGGER IF EXISTS trg_user_roles_assign_at ON user_roles;
DROP TRIGGER IF EXISTS trg_role_permissions_assign_at ON role_permissions;

-- Drop updated_at triggers
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TRIGGER IF EXISTS trg_roles_updated_at ON roles;
DROP TRIGGER IF EXISTS trg_resources_updated_at ON resources;
DROP TRIGGER IF EXISTS trg_actions_updated_at ON actions;
DROP TRIGGER IF EXISTS trg_permissions_updated_at ON permissions;

-- Drop the timestamp function
DROP FUNCTION IF EXISTS set_timestamp();


-- +goose StatementEnd
