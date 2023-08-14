-- name: InsertScheduledInterview :exec
INSERT INTO ScheduledInterview (
  job_seeker_id,
  company_id,
  scheduled_at,
  location,
  notes,
  meeting_link
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: SelectAllScheduledInterviews :many
SELECT * FROM ScheduledInterview;

-- name: SelectScheduledInterviewByID :one
SELECT * FROM ScheduledInterview WHERE id = $1;

-- name: SelectScheduledInterviewByJobSeekerID :many
SELECT * FROM ScheduledInterview WHERE job_seeker_id = $1;

-- name: SelectScheduledInterviewByCompanyID :many
SELECT * FROM ScheduledInterview WHERE company_id = $1;

-- name: UpdateScheduledInterview :exec
UPDATE ScheduledInterview
SET scheduled_at = $1, location = $2, notes = $3, meeting_link = $4
WHERE id = $5;

-- name: DeleteScheduledInterview :exec
DELETE FROM ScheduledInterview WHERE id = $1;
