-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreateUserLocal :one
INSERT INTO users (
  name,
  email,
  password
) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateUserExternal :one
INSERT INTO users (
  name,
  email,
  provider
) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET name = $2, password = $3, provider = $4 WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
