// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: company.sql

package db

import (
	"context"
	"database/sql"
)

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM Company WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCompanyStmt, deleteCompany, id)
	return err
}

const insertCompany = `-- name: InsertCompany :exec
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
) RETURNING id
`

type InsertCompanyParams struct {
	UserID      sql.NullInt32  `json:"user_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Phone       sql.NullString `json:"phone"`
	Website     sql.NullString `json:"website"`
	Logo        sql.NullString `json:"logo"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) InsertCompany(ctx context.Context, arg InsertCompanyParams) error {
	_, err := q.exec(ctx, q.insertCompanyStmt, insertCompany,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Website,
		arg.Logo,
		arg.Description,
	)
	return err
}

const selectAllCompanies = `-- name: SelectAllCompanies :many
SELECT id, user_id, name, email, phone, website, logo, description, created_at, updated_at FROM Company
`

func (q *Queries) SelectAllCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.query(ctx, q.selectAllCompaniesStmt, selectAllCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Website,
			&i.Logo,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectCompanyByID = `-- name: SelectCompanyByID :one
SELECT id, user_id, name, email, phone, website, logo, description, created_at, updated_at FROM Company WHERE id = $1
`

func (q *Queries) SelectCompanyByID(ctx context.Context, id int32) (Company, error) {
	row := q.queryRow(ctx, q.selectCompanyByIDStmt, selectCompanyByID, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Website,
		&i.Logo,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :exec
UPDATE Company SET name = $1, email = $2, phone = $3, website = $4, logo = $5, description = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7
`

type UpdateCompanyParams struct {
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Phone       sql.NullString `json:"phone"`
	Website     sql.NullString `json:"website"`
	Logo        sql.NullString `json:"logo"`
	Description sql.NullString `json:"description"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error {
	_, err := q.exec(ctx, q.updateCompanyStmt, updateCompany,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Website,
		arg.Logo,
		arg.Description,
		arg.ID,
	)
	return err
}
