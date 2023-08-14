-- name: InsertCompany :exec
INSERT INTO Company (
  user_id,
  name,
  email,
  phone,
  website,
  logo,
  description,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7,
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING id;

-- name: SelectAllCompanies :many
SELECT * FROM Company;

-- name: SelectCompanyByID :one
SELECT * FROM Company WHERE id = $1;

-- name: UpdateCompany :exec
UPDATE Company SET name = $1, email = $2, phone = $3, website = $4, logo = $5, description = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7;

-- name: DeleteCompany :exec
DELETE FROM Company WHERE id = $1;
