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
