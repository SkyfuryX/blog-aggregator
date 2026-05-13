-- name: InsertFeed :one
INSERT INTO feeds (name, url, user_id) 
Values (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT 
    f.*,
    u.name as user_name
FROM feeds f
LEFT JOIN users u 
ON f.user_id = u.id;

-- name: GetFeed :one
SELECT * 
FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT 
    f.name as feed_name,
    u.name as user_name
FROM feeds_follows ff
INNER JOIN users u
ON u.id = ff.user_id
INNER JOIN feeds f
ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: CreateFeedFollow :one
WITH inserted_into_feed_follow AS (
    INSERT INTO feeds_follows (user_id, feed_id) 
    VALUES (
        $1,
        $2
    )
    Returning *
)

SELECT 
    i.*,
    f.name as feed_name,
    u.name as user_name
FROM inserted_into_feed_follow i
INNER JOIN users u 
ON i.user_id = u.id
INNER JOIN feeds f 
ON i.feed_id = f.id;
