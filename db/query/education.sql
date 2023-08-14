-- name: InsertEducation :exec
INSERT INTO Education (
  job_seeker_id, institution_id, degree, field_of_study, start_date, end_date, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- name: GetEducations :many
SELECT id, job_seeker_id, institution_id, degree, field_of_study, start_date, end_date, created_at, updated_at FROM Education;

-- name: UpdateEducation :exec
UPDATE Education SET institution_id = $1, degree = $2, field_of_study = $3, start_date = $4, end_date = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6;

-- name: DeleteEducation :exec
DELETE FROM Education WHERE id = $1;
