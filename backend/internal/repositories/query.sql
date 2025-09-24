-- name: GetPosts :many
SELECT post_id, posted_date, address_text, latitude, longtitude, location
FROM posts
WHERE is_active = true
ORDER BY posted_date DESC
LIMIT 50;

-- name: GetPostInfo :one
SELECT post_id, content, incident_date, posted_date, address_text, latitude, longtitude, location
FROM posts
WHERE is_active = true
AND post_id = $1;