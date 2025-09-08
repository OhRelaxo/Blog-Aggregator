-- +goose Up
CREATE TABLE "posts" (id uuid PRIMARY KEY, created_at timestamp NOT NULL, updated_at timestamp NOT NULL, title TEXT not null, url TEXT unique not null, description text not null , published_at timestamp not null, feed_id uuid references feeds(id) NOT NULL);

-- +goose Down
DROP TABLE "posts";