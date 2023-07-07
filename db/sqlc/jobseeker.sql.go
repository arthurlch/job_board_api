// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: jobseeker.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createJobSeeker = `-- name: CreateJobSeeker :one

INSERT INTO JobSeeker (
  user_id,
  resume,
  skills,
  created_at,
  updated_at
) VALUES (
  $1, -- user_id
  $2, -- resume
  $3, -- skills
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id
`

type CreateJobSeekerParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Resume sql.NullString `json:"resume"`
	Skills []string       `json:"skills"`
}

// jobseeker.sql
func (q *Queries) CreateJobSeeker(ctx context.Context, arg CreateJobSeekerParams) (int32, error) {
	row := q.queryRow(ctx, q.createJobSeekerStmt, createJobSeeker, arg.UserID, arg.Resume, pq.Array(arg.Skills))
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteJobSeeker = `-- name: DeleteJobSeeker :exec

DELETE FROM JobSeeker
WHERE id = $1
`

// Specify the condition for updating, such as the "id" column
func (q *Queries) DeleteJobSeeker(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteJobSeekerStmt, deleteJobSeeker, id)
	return err
}

const getJobSeekers = `-- name: GetJobSeekers :many
SELECT id, user_id, resume, skills, created_at, updated_at FROM JobSeeker
`

func (q *Queries) GetJobSeekers(ctx context.Context) ([]Jobseeker, error) {
	rows, err := q.query(ctx, q.getJobSeekersStmt, getJobSeekers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Jobseeker
	for rows.Next() {
		var i Jobseeker
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Resume,
			pq.Array(&i.Skills),
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

const updateJobSeeker = `-- name: UpdateJobSeeker :exec
UPDATE JobSeeker
SET user_id = $1, resume = $2, skills = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4
`

type UpdateJobSeekerParams struct {
	UserID sql.NullInt32  `json:"user_id"`
	Resume sql.NullString `json:"resume"`
	Skills []string       `json:"skills"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdateJobSeeker(ctx context.Context, arg UpdateJobSeekerParams) error {
	_, err := q.exec(ctx, q.updateJobSeekerStmt, updateJobSeeker,
		arg.UserID,
		arg.Resume,
		pq.Array(arg.Skills),
		arg.ID,
	)
	return err
}
