-- +goose Up
CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id INTEGER NOT NULL references feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;