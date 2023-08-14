-- name: InsertExperience :exec
INSERT INTO Experience (
  job_seeker_id, title, company_id, location, start_date, end_date, type_id, description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id;

-- name: SelectAllExperiences :many
SELECT * FROM Experience;

-- name: SelectExperienceByID :one
SELECT * FROM Experience WHERE id = $1;

-- name: UpdateExperience :exec
UPDATE Experience SET title = $1, location = $2, start_date = $3, end_date = $4, type_id = $5, description = $6 WHERE id = $7;

-- name: DeleteExperience :exec
DELETE FROM Experience WHERE id = $1;
