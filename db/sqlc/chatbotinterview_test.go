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

func createRandomChatbotInterviewParams(t *testing.T) InsertChatbotInterviewParams {
	return InsertChatbotInterviewParams{
		JobSeekerID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		JobID:       sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Status:      NullInterviewStatus{"status", true},
		Review:      sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
	}
}

func TestInsertChatbotInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomChatbotInterviewParams(t)
	mock.ExpectExec("INSERT INTO ChatbotInterview").WithArgs(param.JobSeekerID, param.JobID, param.Status, param.Review).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertChatbotInterview(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteChatbotInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM ChatbotInterview WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteChatbotInterview(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllChatbotInterviews(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "job_id", "status", "review", "created_at", "updated_at"}).
		AddRow(1, 1, 2, "status", "review", createdAt, updatedAt)

	mock.ExpectQuery("SELECT (.+) FROM ChatbotInterview").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllChatbotInterviews(context.TODO())
	require.NoError(t, err)
}


func TestSelectChatbotInterviewByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	row := sqlmock.NewRows([]string{"id", "job_seeker_id", "job_id", "status", "review", "created_at", "updated_at"}).
		AddRow(1, 1, 2, "status", "review", createdAt, updatedAt)

	mock.ExpectQuery("SELECT (.+) FROM ChatbotInterview WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectChatbotInterviewByID(context.TODO(), 1)
	require.NoError(t, err)
}


func TestUpdateChatbotInterview(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateChatbotInterviewParams{
		Status: NullInterviewStatus{"status", true},
		Review: sql.NullString{String: "review", Valid: true},
		ID:     int32(1),
	}

	mock.ExpectExec("UPDATE ChatbotInterview SET").WithArgs(param.Status, param.Review, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateChatbotInterview(context.TODO(), param)
	require.NoError(t, err)
}
