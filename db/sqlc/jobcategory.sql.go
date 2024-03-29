// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: jobcategory.sql

package db

import (
	"context"
	"database/sql"
)

const deleteJobCategory = `-- name: DeleteJobCategory :exec
DELETE FROM JobCategory WHERE id = $1
`

func (q *Queries) DeleteJobCategory(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteJobCategoryStmt, deleteJobCategory, id)
	return err
}

const insertJobCategory = `-- name: InsertJobCategory :exec
INSERT INTO JobCategory (name)
VALUES ($1)
RETURNING id
`

func (q *Queries) InsertJobCategory(ctx context.Context, name sql.NullString) error {
	_, err := q.exec(ctx, q.insertJobCategoryStmt, insertJobCategory, name)
	return err
}

const selectAllJobCategories = `-- name: SelectAllJobCategories :many
SELECT id, name FROM JobCategory
`

func (q *Queries) SelectAllJobCategories(ctx context.Context) ([]Jobcategory, error) {
	rows, err := q.query(ctx, q.selectAllJobCategoriesStmt, selectAllJobCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Jobcategory
	for rows.Next() {
		var i Jobcategory
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

const selectJobCategoryByID = `-- name: SelectJobCategoryByID :one
SELECT id, name FROM JobCategory WHERE id = $1
`

func (q *Queries) SelectJobCategoryByID(ctx context.Context, id int32) (Jobcategory, error) {
	row := q.queryRow(ctx, q.selectJobCategoryByIDStmt, selectJobCategoryByID, id)
	var i Jobcategory
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateJobCategory = `-- name: UpdateJobCategory :exec
UPDATE JobCategory SET name = $1 WHERE id = $2
`

type UpdateJobCategoryParams struct {
	Name sql.NullString `json:"name"`
	ID   int32          `json:"id"`
}

func (q *Queries) UpdateJobCategory(ctx context.Context, arg UpdateJobCategoryParams) error {
	_, err := q.exec(ctx, q.updateJobCategoryStmt, updateJobCategory, arg.Name, arg.ID)
	return err
}
