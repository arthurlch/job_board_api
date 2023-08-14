-- jobseeker.sql

-- Insert
-- name: CreateJobSeeker :one
INSERT INTO JobSeeker (
  user_id,
  resume,
  skills,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3::text[],
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- Select all
-- name: GetJobSeekers :many
SELECT id, user_id, resume, skills, created_at, updated_at FROM JobSeeker;

-- Update
-- name: UpdateJobSeeker :exec
UPDATE JobSeeker SET user_id = $1, resume = $2, skills = $3::text[], updated_at = CURRENT_TIMESTAMP WHERE id = $4;

-- Delete
-- name: DeleteJobSeeker :exec
DELETE FROM JobSeeker WHERE id = $1;
