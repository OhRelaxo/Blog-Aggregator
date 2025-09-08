-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select feeds.name, feeds.url, users.name as user_name
from feeds
left join users
on feeds.user_id = users.id;

-- name: GetFeedByURL :one
select * from feeds
where feeds.url = $1;