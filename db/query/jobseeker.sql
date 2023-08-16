-- Insert
-- name: CreateJobSeeker :one
INSERT INTO JobSeeker (
  user_id,
  skills,
  created_at,
  updated_at
) VALUES (
  $1, $2,
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- Select all
-- name: GetJobSeekers :many
SELECT id, user_id, skills, created_at, updated_at FROM JobSeeker;

-- Update
-- name: UpdateJobSeeker :exec
UPDATE JobSeeker SET user_id = $1, skills = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3;

-- Delete
-- name: DeleteJobSeeker :exec
DELETE FROM JobSeeker WHERE id = $1;
