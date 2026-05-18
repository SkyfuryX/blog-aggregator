-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
ON CONFLICT (url) DO UPDATE SET updated_at = now()
RETURNING *;

-- name: GetPostsForUser :many
SELECT p.*, f.name as feed_name
FROM posts p
INNER JOIN feeds_follows ff
ON p.feed_id = ff.feed_id
INNER JOIN feeds f
ON p.feed_id = f.id
WHERE f.user_id = $1 
ORDER BY published_at
Limit $2;