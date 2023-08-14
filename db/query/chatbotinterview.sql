-- name: InsertChatbotInterview :exec
INSERT INTO ChatbotInterview (job_seeker_id, job_id, status, review)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectAllChatbotInterviews :many
SELECT * FROM ChatbotInterview;

-- name: SelectChatbotInterviewByID :one
SELECT * FROM ChatbotInterview WHERE id = $1;

-- name: UpdateChatbotInterview :exec
UPDATE ChatbotInterview SET status = $1, review = $2 WHERE id = $3;

-- name: DeleteChatbotInterview :exec
DELETE FROM ChatbotInterview WHERE id = $1;
