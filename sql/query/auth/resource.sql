-- name: UpsertResourceByID :exec
INSERT INTO resources (resource_id, name)
VALUES ($1, $2)
ON CONFLICT (resource_id) DO UPDATE
SET name = EXCLUDED.name
WHERE resources.name IS DISTINCT FROM EXCLUDED.name;

-- name: UpsertActionByID :exec
INSERT INTO actions (action_id, name, resource_id)
VALUES ($1, $2, $3)
ON CONFLICT (action_id) DO UPDATE
SET name = EXCLUDED.name,
    resource_id = EXCLUDED.resource_id
WHERE actions.name IS DISTINCT FROM EXCLUDED.name
  OR actions.resource_id IS DISTINCT FROM EXCLUDED.resource_id;

-- name: DeleteResourceNotInUse :exec
DELETE FROM resources
WHERE resource_id NOT IN (
  SELECT UNNEST($1::text[])
);

-- name: DeleteActionNotInUse :exec
DELETE FROM actions
WHERE action_id NOT IN (
  SELECT UNNEST($1::text[])
);