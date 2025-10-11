-- name: InsertNewAccount :exec
INSERT INTO users (
    username,
    email,
    password_hash,
    last_login_at
) VALUES (
    $1,
    $2,
    $3,
    $4
); 

-- name: UpdateAccount :exec
UPDATE users
SET 
    username = $2,
    email = $3, 
    password_hash = $4
WHERE user_id = $1;