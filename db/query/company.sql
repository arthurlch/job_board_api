-- company.sql

-- name: CreateCompany :one
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
  $1, -- user_id
  $2, -- name
  $3, -- email
  $4, -- phone
  $5, -- website
  $6, -- logo
  $7, -- description
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;

-- name: GetCompanies :many
SELECT * FROM Company;

-- name: UpdateCompany :exec
UPDATE Company
SET name = $1, email = $2, phone = $3, website = $4, logo = $5, description = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $7; -- Specify the condition for updating, such as the "id" column

-- name: DeleteCompany :exec
DELETE FROM Company
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
