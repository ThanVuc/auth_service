-- name: GetRoles :many
select role_id, name, is_root, is_active, description
from roles
Where
($1::TEXT IS NULL OR $1::TEXT = '' OR name ILIKE '%' || $1::TEXT || '%')
Order by created_at desc
LIMIT NULLIF($2, 0)
OFFSET CASE WHEN $3::INT IS NULL OR $3::INT < 0 THEN 0 ELSE $3::INT END;

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

-- name: IsRootRole :one
select is_root
from roles
where role_id = $1;

-- name: InsertRole :one
insert into roles (name, description)
values ($1, $2)
returning role_id;

-- name: UpdateRole :execrows
update roles
set name = $1, description = $2
where role_id = $3 and is_root = false;

-- name: AddPermissionsToRole :execrows
insert into role_permissions (role_id, perm_id)
select $1, unnest($2::UUID[])
where not exists (
    select 1
    from role_permissions
    where role_id = $1 and perm_id = any($2::UUID[])
);

-- name: RemovePermissionsFromRole :execrows
delete from role_permissions
where role_id = $1 and perm_id = any($2::uuid[]);

-- name: GetPermissionIdsByRole :many
select perm_id
from role_permissions
where role_id = $1;
