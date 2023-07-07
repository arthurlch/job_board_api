-- education.sql

-- name: CreateEducation :one
INSERT INTO Education (
  job_seeker_id,
  institution,
  degree,
  field_of_study,
  start_date,
  end_date,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- institution
  $3, -- degree
  $4, -- field_of_study
  $5, -- start_date
  $6, -- end_date
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetEducations :many
SELECT * FROM Education;

-- name: UpdateEducation :exec
UPDATE Education
SET institution = $1, degree = $2, field_of_study = $3, start_date = $4, end_date = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6; -- Specify the condition for updating, such as the "id" column

-- name: DeleteEducation :exec
DELETE FROM Education
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
