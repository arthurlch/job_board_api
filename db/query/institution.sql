-- name: InsertInstitution :exec
INSERT INTO Institution (name)
VALUES ($1)
RETURNING id;

-- name: SelectAllInstitutions :many
SELECT * FROM Institution;

-- name: SelectInstitutionByID :one
SELECT * FROM Institution WHERE id = $1;

-- name: UpdateInstitution :exec
UPDATE Institution SET name = $1 WHERE id = $2;

-- name: DeleteInstitution :exec
DELETE FROM Institution WHERE id = $1;
