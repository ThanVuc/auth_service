-- +goose Up
-- +goose StatementBegin
ALTER TABLE role_permissions
ADD CONSTRAINT unique_role_perm UNIQUE (role_id, perm_id);

ALTER TABLE permission_actions
ADD CONSTRAINT unique_perm_action UNIQUE (perm_id, action_id);

ALTER TABLE user_roles
ADD CONSTRAINT unique_user_role UNIQUE (user_id, role_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE role_permissions
DROP CONSTRAINT IF EXISTS unique_role_perm;

ALTER TABLE permission_actions
DROP CONSTRAINT IF EXISTS unique_perm_action;

ALTER TABLE user_roles
DROP CONSTRAINT IF EXISTS unique_user_role;
-- +goose StatementEnd
