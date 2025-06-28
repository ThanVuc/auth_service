-- name: UpsertResources :exec
INSERT INTO resources (resource_id, name)
SELECT t1.resource_id, t2.name
FROM unnest($1::TEXT[]) WITH ORDINALITY AS t1(resource_id, ord)
JOIN unnest($2::TEXT[]) WITH ORDINALITY AS t2(name, ord) USING (ord)
ON CONFLICT (resource_id) DO UPDATE
SET name = EXCLUDED.name;

-- name: UpsertActions :exec
INSERT INTO actions (action_id, resource_id, name)
SELECT t1.action_id, t2.resource_id, t3.name
FROM unnest($1::TEXT[]) WITH ORDINALITY AS t1(action_id, ord)
JOIN unnest($2::TEXT[]) WITH ORDINALITY AS t2(resource_id, ord) USING (ord)
JOIN unnest($3::TEXT[]) WITH ORDINALITY AS t3(name, ord) USING (ord)
ON CONFLICT (action_id) DO UPDATE
SET resource_id = EXCLUDED.resource_id,
    name = EXCLUDED.name;

-- name: RemoveOldResources :exec
DELETE FROM resources
WHERE resource_id NOT IN (
  SELECT resource_id FROM unnest($1::TEXT[]) AS resource_id
);

-- name: RemoveOldActions :exec
DELETE FROM actions
WHERE action_id NOT IN (
  SELECT action_id FROM unnest($1::TEXT[]) AS action_id
);


