// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: scheduledinterview.sql

package db

import (
	"context"
	"database/sql"
)

const deleteScheduledInterview = `-- name: DeleteScheduledInterview :exec
DELETE FROM ScheduledInterview WHERE id = $1
`

func (q *Queries) DeleteScheduledInterview(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteScheduledInterviewStmt, deleteScheduledInterview, id)
	return err
}

const insertScheduledInterview = `-- name: InsertScheduledInterview :exec
INSERT INTO ScheduledInterview (
  job_seeker_id,
  company_id,
  scheduled_at,
  location,
  notes,
  meeting_link
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id
`

type InsertScheduledInterviewParams struct {
	JobSeekerID sql.NullInt32  `json:"job_seeker_id"`
	CompanyID   sql.NullInt32  `json:"company_id"`
	ScheduledAt sql.NullTime   `json:"scheduled_at"`
	Location    sql.NullString `json:"location"`
	Notes       sql.NullString `json:"notes"`
	MeetingLink sql.NullString `json:"meeting_link"`
}

func (q *Queries) InsertScheduledInterview(ctx context.Context, arg InsertScheduledInterviewParams) error {
	_, err := q.exec(ctx, q.insertScheduledInterviewStmt, insertScheduledInterview,
		arg.JobSeekerID,
		arg.CompanyID,
		arg.ScheduledAt,
		arg.Location,
		arg.Notes,
		arg.MeetingLink,
	)
	return err
}

const selectAllScheduledInterviews = `-- name: SelectAllScheduledInterviews :many
SELECT id, job_seeker_id, company_id, scheduled_at, location, notes, meeting_link, created_at FROM ScheduledInterview
`

func (q *Queries) SelectAllScheduledInterviews(ctx context.Context) ([]Scheduledinterview, error) {
	rows, err := q.query(ctx, q.selectAllScheduledInterviewsStmt, selectAllScheduledInterviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scheduledinterview
	for rows.Next() {
		var i Scheduledinterview
		if err := rows.Scan(
			&i.ID,
			&i.JobSeekerID,
			&i.CompanyID,
			&i.ScheduledAt,
			&i.Location,
			&i.Notes,
			&i.MeetingLink,
			&i.CreatedAt,
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

const selectScheduledInterviewByCompanyID = `-- name: SelectScheduledInterviewByCompanyID :many
SELECT id, job_seeker_id, company_id, scheduled_at, location, notes, meeting_link, created_at FROM ScheduledInterview WHERE company_id = $1
`

func (q *Queries) SelectScheduledInterviewByCompanyID(ctx context.Context, companyID sql.NullInt32) ([]Scheduledinterview, error) {
	rows, err := q.query(ctx, q.selectScheduledInterviewByCompanyIDStmt, selectScheduledInterviewByCompanyID, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scheduledinterview
	for rows.Next() {
		var i Scheduledinterview
		if err := rows.Scan(
			&i.ID,
			&i.JobSeekerID,
			&i.CompanyID,
			&i.ScheduledAt,
			&i.Location,
			&i.Notes,
			&i.MeetingLink,
			&i.CreatedAt,
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

const selectScheduledInterviewByID = `-- name: SelectScheduledInterviewByID :one
SELECT id, job_seeker_id, company_id, scheduled_at, location, notes, meeting_link, created_at FROM ScheduledInterview WHERE id = $1
`

func (q *Queries) SelectScheduledInterviewByID(ctx context.Context, id int32) (Scheduledinterview, error) {
	row := q.queryRow(ctx, q.selectScheduledInterviewByIDStmt, selectScheduledInterviewByID, id)
	var i Scheduledinterview
	err := row.Scan(
		&i.ID,
		&i.JobSeekerID,
		&i.CompanyID,
		&i.ScheduledAt,
		&i.Location,
		&i.Notes,
		&i.MeetingLink,
		&i.CreatedAt,
	)
	return i, err
}

const selectScheduledInterviewByJobSeekerID = `-- name: SelectScheduledInterviewByJobSeekerID :many
SELECT id, job_seeker_id, company_id, scheduled_at, location, notes, meeting_link, created_at FROM ScheduledInterview WHERE job_seeker_id = $1
`

func (q *Queries) SelectScheduledInterviewByJobSeekerID(ctx context.Context, jobSeekerID sql.NullInt32) ([]Scheduledinterview, error) {
	rows, err := q.query(ctx, q.selectScheduledInterviewByJobSeekerIDStmt, selectScheduledInterviewByJobSeekerID, jobSeekerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scheduledinterview
	for rows.Next() {
		var i Scheduledinterview
		if err := rows.Scan(
			&i.ID,
			&i.JobSeekerID,
			&i.CompanyID,
			&i.ScheduledAt,
			&i.Location,
			&i.Notes,
			&i.MeetingLink,
			&i.CreatedAt,
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

const updateScheduledInterview = `-- name: UpdateScheduledInterview :exec
UPDATE ScheduledInterview
SET scheduled_at = $1, location = $2, notes = $3, meeting_link = $4
WHERE id = $5
`

type UpdateScheduledInterviewParams struct {
	ScheduledAt sql.NullTime   `json:"scheduled_at"`
	Location    sql.NullString `json:"location"`
	Notes       sql.NullString `json:"notes"`
	MeetingLink sql.NullString `json:"meeting_link"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateScheduledInterview(ctx context.Context, arg UpdateScheduledInterviewParams) error {
	_, err := q.exec(ctx, q.updateScheduledInterviewStmt, updateScheduledInterview,
		arg.ScheduledAt,
		arg.Location,
		arg.Notes,
		arg.MeetingLink,
		arg.ID,
	)
	return err
}
