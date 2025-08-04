-- name: GetResources :many
select resource_id, name
from resources;

-- name: GetActions :many
select action_id, name
from actions
where resource_id = $1;

-- name: GetPermissions :many
select perm_id, name, is_root, description
from permissions
Where
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%') AND
($2::TEXT IS NULL OR $2::TEXT = '' OR resource_id = $2::TEXT)
Order by perm_id
limit $3 offset $4;

-- name: CountTotalPermissions :one
select count(perm_id) as total
from permissions
Where
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%') AND
($2::TEXT IS NULL OR $2::TEXT = '' OR resource_id = $2::TEXT);

-- name: CountRootPermissions :one
select count(perm_id) as total
from permissions
where is_root = true and
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%') AND
($2::TEXT IS NULL OR $2::TEXT = '' OR resource_id = $2::TEXT);

-- name: InsertPermission :one
INSERT INTO permissions (name, resource_id, description)
VALUES ($1, $2, $3)
RETURNING perm_id;

-- name: UpdatePermission :one
UPDATE permissions
SET name = $2,
    resource_id = $3,
    description = $4
WHERE perm_id = $1
RETURNING perm_id;

-- name: GetActionsByPermissionId :many
select pa.pa_id, pa.action_id
from permission_actions pa
where pa.perm_id = $1;

-- name: AddActionToPermission :exec
INSERT INTO permission_actions (perm_id, action_id)
SELECT $1, unnest($2::TEXT[])
WHERE NOT EXISTS (
    SELECT 1
    FROM permission_actions
    WHERE perm_id = $1 AND action_id = ANY($2::TEXT[])
);

-- name: DeleteActionToPermission :exec
DELETE FROM permission_actions
WHERE perm_id = $1 AND action_id = ANY($2::TEXT[]);

-- name: GetPermission :many
SELECT
    p.perm_id,
    p.name AS permission_name,
    p.resource_id,
    p.is_root,
    r.name AS resource_name,
    p.description,
    p.updated_at,
    p.created_at,
    a.action_id,
    a.name AS action_name
FROM permissions p
JOIN resources r ON r.resource_id = p.resource_id
LEFT JOIN permission_actions pa ON pa.perm_id = p.perm_id
LEFT JOIN actions a ON a.action_id = pa.action_id
WHERE p.perm_id = $1;

-- name: DeletePermission :exec
delete from permissions
where perm_id = $1;
