-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  url TEXT NOT NULL UNIQUE,
  description TEXT NOT NULL,
  published_at TIMESTAMP,
  feed_id UUID NOT NULL,
  FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
