// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: experience.sql

package db

import (
	"context"
	"database/sql"
)

const deleteExperience = `-- name: DeleteExperience :exec
DELETE FROM Experience WHERE id = $1
`

func (q *Queries) DeleteExperience(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteExperienceStmt, deleteExperience, id)
	return err
}

const insertExperience = `-- name: InsertExperience :exec
INSERT INTO Experience (
  job_seeker_id, title, company_id, location, start_date, end_date, type_id, description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id
`

type InsertExperienceParams struct {
	JobSeekerID sql.NullInt32  `json:"job_seeker_id"`
	Title       sql.NullString `json:"title"`
	CompanyID   sql.NullInt32  `json:"company_id"`
	Location    sql.NullString `json:"location"`
	StartDate   sql.NullTime   `json:"start_date"`
	EndDate     sql.NullTime   `json:"end_date"`
	TypeID      sql.NullInt32  `json:"type_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) InsertExperience(ctx context.Context, arg InsertExperienceParams) error {
	_, err := q.exec(ctx, q.insertExperienceStmt, insertExperience,
		arg.JobSeekerID,
		arg.Title,
		arg.CompanyID,
		arg.Location,
		arg.StartDate,
		arg.EndDate,
		arg.TypeID,
		arg.Description,
	)
	return err
}

const selectAllExperiences = `-- name: SelectAllExperiences :many
SELECT id, job_seeker_id, title, company_id, location, start_date, end_date, type_id, description, created_at, updated_at FROM Experience
`

func (q *Queries) SelectAllExperiences(ctx context.Context) ([]Experience, error) {
	rows, err := q.query(ctx, q.selectAllExperiencesStmt, selectAllExperiences)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Experience
	for rows.Next() {
		var i Experience
		if err := rows.Scan(
			&i.ID,
			&i.JobSeekerID,
			&i.Title,
			&i.CompanyID,
			&i.Location,
			&i.StartDate,
			&i.EndDate,
			&i.TypeID,
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

const selectExperienceByID = `-- name: SelectExperienceByID :one
SELECT id, job_seeker_id, title, company_id, location, start_date, end_date, type_id, description, created_at, updated_at FROM Experience WHERE id = $1
`

func (q *Queries) SelectExperienceByID(ctx context.Context, id int32) (Experience, error) {
	row := q.queryRow(ctx, q.selectExperienceByIDStmt, selectExperienceByID, id)
	var i Experience
	err := row.Scan(
		&i.ID,
		&i.JobSeekerID,
		&i.Title,
		&i.CompanyID,
		&i.Location,
		&i.StartDate,
		&i.EndDate,
		&i.TypeID,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateExperience = `-- name: UpdateExperience :exec
UPDATE Experience SET title = $1, location = $2, start_date = $3, end_date = $4, type_id = $5, description = $6 WHERE id = $7
`

type UpdateExperienceParams struct {
	Title       sql.NullString `json:"title"`
	Location    sql.NullString `json:"location"`
	StartDate   sql.NullTime   `json:"start_date"`
	EndDate     sql.NullTime   `json:"end_date"`
	TypeID      sql.NullInt32  `json:"type_id"`
	Description sql.NullString `json:"description"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateExperience(ctx context.Context, arg UpdateExperienceParams) error {
	_, err := q.exec(ctx, q.updateExperienceStmt, updateExperience,
		arg.Title,
		arg.Location,
		arg.StartDate,
		arg.EndDate,
		arg.TypeID,
		arg.Description,
		arg.ID,
	)
	return err
}
