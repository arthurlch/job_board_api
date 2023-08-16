-- Insert
-- name: CreateJobSeeker :one
INSERT INTO JobSeeker (
  user_id,
  created_at,
  updated_at
) VALUES (
  $1,
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- Select all
-- name: GetJobSeekers :many
SELECT id, user_id, created_at, updated_at FROM JobSeeker;

-- Update
-- name: UpdateJobSeeker :exec
UPDATE JobSeeker SET user_id = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2;

-- Delete
-- name: DeleteJobSeeker :exec
DELETE FROM JobSeeker WHERE id = $1;
