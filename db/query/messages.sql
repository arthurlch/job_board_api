-- name: InsertMessage :exec
INSERT INTO Messages (sender_id, receiver_id, content, sender_type)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectMessagesBySenderAndReceiver :many
SELECT * FROM Messages WHERE sender_id = $1 AND receiver_id = $2;

-- name: DeleteMessage :exec
DELETE FROM Messages WHERE id = $1;
