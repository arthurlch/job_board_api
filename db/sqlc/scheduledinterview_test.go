package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomScheduledInterviewParams(t *testing.T) InsertScheduledInterviewParams {
	return InsertScheduledInterviewParams{
		JobSeekerID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		CompanyID:   sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		ScheduledAt: sql.NullTime{Time: gofakeit.Date(), Valid: true},
		Location:    sql.NullString{String: gofakeit.City(), Valid: true},
		Notes:       sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		MeetingLink: sql.NullString{String: gofakeit.URL(), Valid: true},
	}
}

func TestInsertScheduledInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomScheduledInterviewParams(t)
	mock.ExpectExec("INSERT INTO ScheduledInterview").WithArgs(param.JobSeekerID, param.CompanyID, param.ScheduledAt, param.Location, param.Notes, param.MeetingLink).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertScheduledInterview(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteScheduledInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM ScheduledInterview WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteScheduledInterview(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllScheduledInterviews(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "company_id", "scheduled_at", "location", "notes", "meeting_link", "created_at"}).
		AddRow(1, 1, 2, time.Now(), "location", "notes", "meeting_link", time.Now())

	mock.ExpectQuery("SELECT (.+) FROM ScheduledInterview").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllScheduledInterviews(context.TODO())
	require.NoError(t, err)
}


func TestUpdateScheduledInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateScheduledInterviewParams{
		ScheduledAt: sql.NullTime{Time: gofakeit.Date(), Valid: true},
		Location:    sql.NullString{String: "location", Valid: true},
		Notes:       sql.NullString{String: "notes", Valid: true},
		MeetingLink: sql.NullString{String: "meeting_link", Valid: true},
		ID:          int32(1),
	}

	mock.ExpectExec("UPDATE ScheduledInterview SET").WithArgs(param.ScheduledAt, param.Location, param.Notes, param.MeetingLink, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateScheduledInterview(context.TODO(), param)
	require.NoError(t, err)
}
