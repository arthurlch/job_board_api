-- job.sql

-- Insert
-- name: CreateJob :one
INSERT INTO Job (
  title,
  description,
  requirements,
  location,
  salary,
  company_id,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6,
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- Select all
-- name: GetJobs :many
SELECT id, title, description, requirements, location, salary, company_id, created_at, updated_at FROM Job;

-- Update
-- name: UpdateJob :exec
UPDATE Job SET title = $1, description = $2, requirements = $3, location = $4, salary = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6;

-- Delete
-- name: DeleteJob :exec
DELETE FROM Job WHERE id = $1;
