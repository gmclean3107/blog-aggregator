-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (
        id,
        created_at,
        updated_at,
        user_id,
        feed_id
    )
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, users.name AS "user_name", feeds.name AS "feed_name"
FROM inserted_feed_follow
JOIN users ON inserted_feed_follow.user_id = users.id
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT users.name AS "user_name", feeds.name AS "feed_name"
FROM feed_follows
JOIN feeds ON feeds.id = feed_follows.feed_id
JOIN users ON users.id = feed_follows.user_id
WHERE $1 = feed_follows.user_id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE $1 = user_id AND $2 = feed_id;