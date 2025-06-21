-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column() -- update_at trigger function
RETURNS TRIGGER AS $$
BEGIN
  NEW.update_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_assign_at_column() -- assign_at trigger function
RETURNS TRIGGER AS $$
BEGIN
  NEW.assign_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggers for updating timestamps
CREATE TRIGGER trg_users_update_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trg_roles_update_at
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trg_permissions_update_at
BEFORE UPDATE ON permissions
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Triggers for updating assign_at timestamps
CREATE TRIGGER trg_user_roles_assign_at
BEFORE UPDATE ON user_roles
FOR EACH ROW
EXECUTE FUNCTION update_assign_at_column();

CREATE TRIGGER trg_role_permissions_assign_at
BEFORE UPDATE ON role_permissions
FOR EACH ROW
EXECUTE FUNCTION update_assign_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_users_update_at ON users;
DROP TRIGGER IF EXISTS trg_roles_update_at ON roles;
DROP TRIGGER IF EXISTS trg_permissions_update_at ON permissions;
DROP TRIGGER IF EXISTS trg_user_roles_assign_at ON user_roles;
DROP TRIGGER IF EXISTS trg_role_permissions_assign_at ON role_permissions;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP FUNCTION IF EXISTS update_assign_at_column();
-- +goose StatementEnd
