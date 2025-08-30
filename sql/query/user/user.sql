-- name: GetUsers :many
SELECT user_id, email, lock_end, lock_reason, last_login_at
FROM users
WHERE ($1::TEXT IS NULL OR $1::TEXT = '' OR email ILIKE '%' || $1::TEXT || '%')
ORDER BY created_at DESC
LIMIT NULLIF($2, 0)
OFFSET CASE WHEN $3::INT IS NULL OR $3::INT < 0 THEN 0 ELSE $3::INT END;

-- name: CountTotalUsers :one
SELECT count(user_id) as total
FROM users
WHERE ($1::TEXT IS NULL OR $1::TEXT = '' OR email ILIKE '%' || $1::TEXT || '%');

-- name: GetRoleIDsByUserID :many
SELECT ur.role_id
FROM user_roles ur
WHERE ur.user_id = $1;

-- name: AddNewRolesToUser :exec
insert into user_roles (user_id, role_id)
select $1, unnest($2::UUID[])
where not exists (
    select 1
    from user_roles
    where user_id = $1 and role_id = any($2::UUID[])
);

-- name: RemoveRolesFromUser :exec
delete from user_roles
where user_id = $1 and role_id = any($2::uuid[]);

-- name: LoginWithExternalProvider :one
select u.user_id, u.email, u.lock_end, u.lock_reason
from users u
join external_provider as ep on ep.user_id = u.user_id
WHERE u.email = $1 AND ep.sub = $2;

-- name: UpdateUserLastLogin :exec
update users
set last_login_at = now()
where user_id = $1;

-- name: GetUser :many
SELECT
    u.user_id,
    u.email,
    u.lock_end,
    u.lock_reason,
    u.created_at,
    u.updated_at,
    u.last_login_at,
    r.role_id,
    r.name AS role_name,
    r.description AS role_description
FROM users u
LEFT JOIN user_roles ur ON ur.user_id = u.user_id
LEFT JOIN roles r ON r.role_id = ur.role_id
WHERE u.user_id = $1;


-- name: LockUser :exec
UPDATE users 
SET lock_end = '9999-12-31', lock_reason = $2, updated_at = NOW()
WHERE user_id = $1;


-- name: UnlockUser :exec
UPDATE users
SET lock_end = NOW(), updated_at = NOW()
WHERE user_id = $1;

-- name: GetLockEndByUserID :one
SELECT lock_end
FROM users
WHERE user_id = $1;