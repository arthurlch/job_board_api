-- user.sql

-- name: CreateUser :one
INSERT INTO "User" (
  name,
  email,
  phone,
  created_at,
  updated_at
) VALUES (
  $1, -- name
  $2, -- email
  $3, -- phone
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetUsers :many
SELECT * FROM "User";

-- name: UpdateUser :exec
UPDATE "User"
SET name = $1, email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4; -- Specify the condition for updating, such as the "id" column

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
