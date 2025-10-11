-- name: SelectPosts :many
SELECT post_id, posted_date, address_text, latitude, longtitude, location
FROM posts
WHERE is_active = true
ORDER BY posted_date DESC
LIMIT 50;

-- name: SelectPostInfo :one
SELECT post_id, content, incident_date, posted_date, address_text, latitude, longtitude, location
FROM posts
WHERE is_active = true
AND post_id = $1;

-- name: InsertNewPost :exec
INSERT INTO posts (
    content,
    incident_date,
    latitude,
    longtitude,
    address_text,
    location
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    ST_SetSRID(ST_MakePoint($4, $3), 4326)
);


