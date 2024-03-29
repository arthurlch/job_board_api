// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: companyentity.sql

package db

import (
	"context"
	"database/sql"
)

const deleteCompanyEntity = `-- name: DeleteCompanyEntity :exec
DELETE FROM CompanyEntity WHERE id = $1
`

func (q *Queries) DeleteCompanyEntity(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCompanyEntityStmt, deleteCompanyEntity, id)
	return err
}

const insertCompanyEntity = `-- name: InsertCompanyEntity :exec
INSERT INTO CompanyEntity (name)
VALUES ($1)
RETURNING id
`

func (q *Queries) InsertCompanyEntity(ctx context.Context, name sql.NullString) error {
	_, err := q.exec(ctx, q.insertCompanyEntityStmt, insertCompanyEntity, name)
	return err
}

const selectAllCompanyEntities = `-- name: SelectAllCompanyEntities :many
SELECT id, name FROM CompanyEntity
`

func (q *Queries) SelectAllCompanyEntities(ctx context.Context) ([]Companyentity, error) {
	rows, err := q.query(ctx, q.selectAllCompanyEntitiesStmt, selectAllCompanyEntities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Companyentity
	for rows.Next() {
		var i Companyentity
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const selectCompanyEntityByID = `-- name: SelectCompanyEntityByID :one
SELECT id, name FROM CompanyEntity WHERE id = $1
`

func (q *Queries) SelectCompanyEntityByID(ctx context.Context, id int32) (Companyentity, error) {
	row := q.queryRow(ctx, q.selectCompanyEntityByIDStmt, selectCompanyEntityByID, id)
	var i Companyentity
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateCompanyEntity = `-- name: UpdateCompanyEntity :exec
UPDATE CompanyEntity SET name = $1 WHERE id = $2
`

type UpdateCompanyEntityParams struct {
	Name sql.NullString `json:"name"`
	ID   int32          `json:"id"`
}

func (q *Queries) UpdateCompanyEntity(ctx context.Context, arg UpdateCompanyEntityParams) error {
	_, err := q.exec(ctx, q.updateCompanyEntityStmt, updateCompanyEntity, arg.Name, arg.ID)
	return err
}
