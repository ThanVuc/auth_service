-- name: InsertUser :one
INSERT INTO users (email, password_hash, last_login_at)
VALUES ($1, $2, $3)
RETURNING user_id, email, created_at, updated_at;

-- name: InsertExternalProvider :one
INSERT INTO external_provider (sub, provider, user_id)
VALUES ($1, $2, $3)
RETURNING sub;

-- name: HasPermission :one
SELECT EXISTS (
    SELECT 1
    FROM role_permissions rp
    JOIN permissions p ON p.perm_id = rp.perm_id
    JOIN resources rc ON rc.resource_id = p.resource_id
    JOIN permission_actions pa ON pa.perm_id = p.perm_id
    JOIN actions at ON at.action_id = pa.action_id
    WHERE rp.role_id = ANY($1::uuid[])
      AND rc.name = $2
      AND at.name = $3
) AS has_permission;

-- name: GetUserAuthInfo :many
SELECT
    rp.role_id,
    r.name          AS role_name,
    p.perm_id,
    p.name          AS perm_name,
    rc.resource_id,
    rc.name         AS resource_name,
    at.action_id,
    at.name         AS action_name
FROM role_permissions rp
JOIN roles r ON r.role_id = rp.role_id
JOIN permissions p ON p.perm_id = rp.perm_id
JOIN resources rc ON rc.resource_id = p.resource_id
JOIN permission_actions pa ON pa.perm_id = p.perm_id
JOIN actions at ON at.action_id = pa.action_id
WHERE rp.role_id = ANY($1::uuid[]);

