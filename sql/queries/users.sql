-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
values ($1, $2, $3, $4)
returning *;

-- name: GetUser :one
select *
from users
where name = $1;

-- name: ResetUsers :exec
DELETE FROM users;

-- name: GetUsers :many
select users.name from users;