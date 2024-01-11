-- name: create_author :one
INSERT INTO authors ( 
    name, bio, email
) VALUES (
    $1, $2, $3
)
RETURNING *;