-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id
INNER JOIN users on inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT *, users.name as user_name, feeds.name as feed_name FROM feed_follows
LEFT JOIN feeds on feed_follows.feed_id = feeds.id
LEFT JOIN users on feed_follows.user_id = users.id
where users.name = $1;

-- name: DeleteFollow :exec
DELETE FROM feed_follows
where user_id = $1
and feed_id = $2;
