-- experience.sql

-- name: CreateExperience :one
INSERT INTO Experience (
  job_seeker_id,
  title,
  company,
  location,
  start_date,
  end_date,
  description,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- title
  $3, -- company
  $4, -- location
  $5, -- start_date
  $6, -- end_date
  $7, -- description
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetExperiences :many
SELECT * FROM Experience;

-- name: UpdateExperience :exec
UPDATE Experience
SET title = $1, company = $2, location = $3, start_date = $4, end_date = $5, description = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $7; -- Specify the condition for updating, such as the "id" column

-- name: DeleteExperience :exec
DELETE FROM Experience
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
