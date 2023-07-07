-- job.sql

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
  $1, -- title
  $2, -- description
  $3, -- requirements
  $4, -- location
  $5, -- salary
  $6, -- company_id
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetJobs :many
SELECT * FROM Job;

-- name: UpdateJob :exec
UPDATE Job
SET title = $1, description = $2, requirements = $3, location = $4, salary = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6; -- Specify the condition for updating, such as the "id" column

-- name: DeleteJob :exec
DELETE FROM Job
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
