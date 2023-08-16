// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: institution.sql

package db

import (
	"context"
	"database/sql"
)

const deleteInstitution = `-- name: DeleteInstitution :exec
DELETE FROM Institution WHERE id = $1
`

func (q *Queries) DeleteInstitution(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteInstitutionStmt, deleteInstitution, id)
	return err
}

const insertInstitution = `-- name: InsertInstitution :exec
INSERT INTO Institution (name)
VALUES ($1)
RETURNING id
`

func (q *Queries) InsertInstitution(ctx context.Context, name sql.NullString) error {
	_, err := q.exec(ctx, q.insertInstitutionStmt, insertInstitution, name)
	return err
}

const selectAllInstitutions = `-- name: SelectAllInstitutions :many
SELECT id, name FROM Institution
`

func (q *Queries) SelectAllInstitutions(ctx context.Context) ([]Institution, error) {
	rows, err := q.query(ctx, q.selectAllInstitutionsStmt, selectAllInstitutions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Institution
	for rows.Next() {
		var i Institution
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

const selectInstitutionByID = `-- name: SelectInstitutionByID :one
SELECT id, name FROM Institution WHERE id = $1
`

func (q *Queries) SelectInstitutionByID(ctx context.Context, id int32) (Institution, error) {
	row := q.queryRow(ctx, q.selectInstitutionByIDStmt, selectInstitutionByID, id)
	var i Institution
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateInstitution = `-- name: UpdateInstitution :exec
UPDATE Institution SET name = $1 WHERE id = $2
`

type UpdateInstitutionParams struct {
	Name sql.NullString `json:"name"`
	ID   int32          `json:"id"`
}

func (q *Queries) UpdateInstitution(ctx context.Context, arg UpdateInstitutionParams) error {
	_, err := q.exec(ctx, q.updateInstitutionStmt, updateInstitution, arg.Name, arg.ID)
	return err
}