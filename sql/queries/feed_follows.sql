-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id) 
  VALUES (
    $1,  -- id (UUID)
    $2,  -- created_at (timestamp)
    $3,  -- updated_at (timestamp)
    $4,  -- user_id (UUID)
    $5   -- feed_id (UUID)
  )
  RETURNING *
)
SELECT
    iff.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow iff
JOIN users ON iff.user_id = users.id
JOIN feeds ON iff.feed_id = feeds.id;

-- name: DeleteFeedFollow :exec
WITH deleted_feeds AS (
    DELETE FROM feeds 
    WHERE url = $1
    RETURNING id
)
DELETE FROM feed_follows
WHERE feed_id IN (SELECT id FROM deleted_feeds);

-- name: GetFeedFollowsForUser :many
SELECT 
  feed_follows.*,
  feeds.name AS feed_name,
  users.name AS user_name
  FROM feed_follows 
  JOIN users ON feed_follows.user_id = users.id
  JOIN feeds ON feed_follows.feed_id = feeds.id 
  WHERE feed_follows.user_id = $1;
