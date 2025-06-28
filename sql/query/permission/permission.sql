-- name: GetResources :many
select resource_id, name
from resources;

-- name: GetActions :many
select action_id, name
from actions
where resource_id = $1;