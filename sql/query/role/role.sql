-- name: GetRoles :many
select role_id, name, is_root, is_active, description
from roles
Where
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%')
Order by role_id
limit $2 offset $3;

-- name: GetRoleById :many
select 
    r.role_id as role_id,
    r.name as role_name,
    r.description,
    r.is_root,
    r.created_at,
    r.updated_at,
    r.is_active,
    p.perm_id as permission_id,
    p.name as permission_name,
    p.description as permission_description
from roles r
left join role_permissions rp on r.role_id = rp.role_id
left join permissions p on rp.perm_id = p.perm_id
where r.role_id = $1;

-- name: CountTotalRoles :one
select count(role_id) as total
from roles
Where
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%');

-- name: CountRootRoles :one
select count(role_id) as total
from roles
where is_root = true and
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%');

-- name: DisableOrEnableRole :execrows
UPDATE roles
SET is_active = NOT is_active
WHERE role_id = $1 and is_root = false;

-- name: DeleteRole :execrows
DELETE FROM roles
WHERE role_id = $1 and is_root = false;

-- name: CountUsersByRoles :many
select r.role_id, count(ur.user_id) as total_users
from roles r
left join user_roles ur on r.role_id = ur.role_id
where r.role_id = any($1::uuid[])
group by r.role_id;
