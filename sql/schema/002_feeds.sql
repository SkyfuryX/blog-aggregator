-- +goose Up
CREATE TABLE feeds (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
url TEXT UNIQUE NOT NULL,
user_id UUID NOT NULL references users(id) ON DELETE CASCADE,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feeds;