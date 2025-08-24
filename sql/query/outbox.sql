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
