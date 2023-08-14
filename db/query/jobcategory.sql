-- name: InsertJobCategory :exec
INSERT INTO JobCategory (name)
VALUES ($1)
RETURNING id;

-- name: SelectAllJobCategories :many
SELECT * FROM JobCategory;

-- name: SelectJobCategoryByID :one
SELECT * FROM JobCategory WHERE id = $1;

-- name: UpdateJobCategory :exec
UPDATE JobCategory SET name = $1 WHERE id = $2;

-- name: DeleteJobCategory :exec
DELETE FROM JobCategory WHERE id = $1;