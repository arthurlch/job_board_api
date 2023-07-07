-- jobseeker.sql

-- name: CreateJobSeeker :one
INSERT INTO JobSeeker (
  user_id,
  resume,
  skills,
  created_at,
  updated_at
) VALUES (
  $1, -- user_id
  $2, -- resume
  $3, -- skills
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetJobSeekers :many
SELECT * FROM JobSeeker;

-- name: UpdateJobSeeker :exec
UPDATE JobSeeker
SET user_id = $1, resume = $2, skills = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4; -- Specify the condition for updating, such as the "id" column

-- name: DeleteJobSeeker :exec
DELETE FROM JobSeeker
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
