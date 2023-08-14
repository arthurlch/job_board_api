-- name: InsertJobView :exec
INSERT INTO JobViews (job_id, view_count)
VALUES ($1, $2);

-- name: SelectJobViewsByJobID :one
SELECT * FROM JobViews WHERE job_id = $1;

-- name: UpdateJobViews :exec
UPDATE JobViews SET view_count = $1 WHERE job_id = $2;
