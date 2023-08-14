-- name: InsertApplication :exec
INSERT INTO Application (
  job_seeker_id, job_id, cover_letter, resume, status
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id;

-- name: SelectAllApplications :many
SELECT * FROM Application;

-- name: SelectApplicationByID :one
SELECT * FROM Application WHERE id = $1;

-- name: UpdateApplication :exec
UPDATE Application SET cover_letter = $1, resume = $2, status = $3 WHERE id = $4;

-- name: DeleteApplication :exec
DELETE FROM Application WHERE id = $1;
