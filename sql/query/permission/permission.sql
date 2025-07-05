-- name: GetResources :many
select resource_id, name
from resources;

-- name: GetActions :many
select action_id, name
from actions
where resource_id = $1;

-- name: GetPermissions :many
select perm_id, name, is_root
from permissions
Where
($1::TEXT is null or name ilike '%' || $1 || '%') and
($2::TEXT is null or resource_id = $1)
Order by perm_id
limit $3 offset $4;

-- name: CountTotalPermissions :one
select count(perm_id) as total
from permissions
Where
($1::TEXT is null or name ilike '%' || $1 || '%') and
($2::TEXT is null or resource_id = $1);

-- name: CountRootPermissions :one
select count(perm_id) as total
from permissions
where is_root = true and
($1::TEXT is null or name ilike '%' || $1 || '%') and
($2::TEXT is null or resource_id = $2);

-- name: DeletePermission :exec
delete from permissions
where perm_id = $1;

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