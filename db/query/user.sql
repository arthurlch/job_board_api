-- name: InsertUser :exec
INSERT INTO "User" (name, email, phone, role) VALUES ($1, $2, $3, $4) RETURNING id;
-- name: SelectUsers :many
SELECT * FROM "User";
-- name: SelectUserByID :one
SELECT * FROM "User" WHERE id = $1;
-- name: UpdateUser :exec
UPDATE "User" SET name = $1, email = $2, phone = $3, role = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5;
-- name: DeleteUser :exec
DELETE FROM "User" WHERE id = $1;
