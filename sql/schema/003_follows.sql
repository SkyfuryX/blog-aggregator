-- +goose Up
CREATE TABLE feeds_follows (
id SERIAL PRIMARY KEY,
created_at TIMESTAMP DEFAULT now() NOT NULL,
updated_at TIMESTAMP DEFAULT now() NOT NULL,
user_id UUID NOT NULL references users(id) ON DELETE CASCADE,
feed_id INTEGER NOT NULL references feeds(id) ON DELETE CASCADE,
UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feeds_follows;