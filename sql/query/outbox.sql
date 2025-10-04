-- name: InsertOutbox :one
INSERT INTO outbox (
    aggregate_type,
    aggregate_id,
    event_type,
    payload,
    request_id
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, occurred_at;

-- name: InsertOutboxBulk :exec
INSERT INTO outbox (
    aggregate_type,
    aggregate_id,
    event_type,
    payload,
    request_id
)
SELECT 
    unnest($1::text[])   AS aggregate_type,
    unnest($2::text[])   AS aggregate_id,
    unnest($3::text[])   AS event_type,
    unnest($4::jsonb[])  AS payload,
    unnest($5::text[])   AS request_id;
