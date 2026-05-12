-- name: InsertFeed :one
INSERT INTO feeds (name, url, user_id) Values (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.*, u.name  FROM feeds f
LEFT JOIN users u 
ON f.user_id = u.id;