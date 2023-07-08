// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: education.sql

package db

import (
	"context"
	"database/sql"
)

const createEducation = `-- name: CreateEducation :one

INSERT INTO Education (
  job_seeker_id,
  institution,
  degree,
  field_of_study,
  start_date,
  end_date,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- institution
  $3, -- degree
  $4, -- field_of_study
  $5, -- start_date
  $6, -- end_date
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id
`

type CreateEducationParams struct {
	JobSeekerID  sql.NullInt32  `json:"job_seeker_id"`
	Institution  sql.NullString `json:"institution"`
	Degree       sql.NullString `json:"degree"`
	FieldOfStudy sql.NullString `json:"field_of_study"`
	StartDate    sql.NullTime   `json:"start_date"`
	EndDate      sql.NullTime   `json:"end_date"`
}

// education.sql
func (q *Queries) CreateEducation(ctx context.Context, arg CreateEducationParams) (int32, error) {
	row := q.queryRow(ctx, q.createEducationStmt, createEducation,
		arg.JobSeekerID,
		arg.Institution,
		arg.Degree,
		arg.FieldOfStudy,
		arg.StartDate,
		arg.EndDate,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteEducation = `-- name: DeleteEducation :exec

DELETE FROM Education
WHERE id = $1
`

// Specify the condition for updating, such as the "id" column
func (q *Queries) DeleteEducation(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteEducationStmt, deleteEducation, id)
	return err
}

const getEducations = `-- name: GetEducations :many
SELECT id, job_seeker_id, institution, degree, field_of_study, start_date, end_date, created_at, updated_at FROM Education
`

func (q *Queries) GetEducations(ctx context.Context) ([]Education, error) {
	rows, err := q.query(ctx, q.getEducationsStmt, getEducations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Education
	for rows.Next() {
		var i Education
		if err := rows.Scan(
			&i.ID,
			&i.JobSeekerID,
			&i.Institution,
			&i.Degree,
			&i.FieldOfStudy,
			&i.StartDate,
			&i.EndDate,
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

const updateEducation = `-- name: UpdateEducation :exec
UPDATE Education
SET institution = $1, degree = $2, field_of_study = $3, start_date = $4, end_date = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6
`

type UpdateEducationParams struct {
	Institution  sql.NullString `json:"institution"`
	Degree       sql.NullString `json:"degree"`
	FieldOfStudy sql.NullString `json:"field_of_study"`
	StartDate    sql.NullTime   `json:"start_date"`
	EndDate      sql.NullTime   `json:"end_date"`
	ID           int32          `json:"id"`
}

func (q *Queries) UpdateEducation(ctx context.Context, arg UpdateEducationParams) error {
	_, err := q.exec(ctx, q.updateEducationStmt, updateEducation,
		arg.Institution,
		arg.Degree,
		arg.FieldOfStudy,
		arg.StartDate,
		arg.EndDate,
		arg.ID,
	)
	return err
}