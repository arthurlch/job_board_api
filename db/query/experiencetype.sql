-- name: InsertExperienceType :exec
INSERT INTO ExperienceType (name)
VALUES ($1)
RETURNING id;

-- name: SelectAllExperienceTypes :many
SELECT * FROM ExperienceType;

-- name: SelectExperienceTypeByID :one
SELECT * FROM ExperienceType WHERE id = $1;

-- name: UpdateExperienceType :exec
UPDATE ExperienceType SET name = $1 WHERE id = $2;

-- name: DeleteExperienceType :exec
DELETE FROM ExperienceType WHERE id = $1;
