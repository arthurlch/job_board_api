-- name: InsertCompanyEntity :exec
INSERT INTO CompanyEntity (name)
VALUES ($1)
RETURNING id;

-- name: SelectAllCompanyEntities :many
SELECT * FROM CompanyEntity;

-- name: SelectCompanyEntityByID :one
SELECT * FROM CompanyEntity WHERE id = $1;

-- name: UpdateCompanyEntity :exec
UPDATE CompanyEntity SET name = $1 WHERE id = $2;

-- name: DeleteCompanyEntity :exec
DELETE FROM CompanyEntity WHERE id = $1;
