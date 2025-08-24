-- name: GetUsers :many
SELECT user_id, email, lock_end, lock_reason
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

