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

func createRandomApplicationParams(t *testing.T) InsertApplicationParams {
	return InsertApplicationParams{
		JobSeekerID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		JobID:       sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		CoverLetter: sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Resume:      sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Status:      NullApplicationStatus{ApplicationStatus: "applied", Valid: true},
	}
}




func TestInsertApplication(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomApplicationParams(t)
	mock.ExpectExec("INSERT INTO Application").WithArgs(param.JobSeekerID, param.JobID, param.CoverLetter, param.Resume, param.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertApplication(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteApplication(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Application WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteApplication(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllApplications(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "job_id", "cover_letter", "resume", "status", "created_at", "updated_at"}).
		AddRow(1, 1, 2, "cover_letter", "resume", "status", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Application").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllApplications(context.TODO())
	require.NoError(t, err)
}

func TestSelectApplicationByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "job_seeker_id", "job_id", "cover_letter", "resume", "status", "created_at", "updated_at"}).
		AddRow(1, 1, 2, "cover_letter", "resume", "status", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Application WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectApplicationByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateApplication(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateApplicationParams{
		CoverLetter: sql.NullString{String: "cover_letter", Valid: true},
		Resume:      sql.NullString{String: "resume", Valid: true},
		Status:      NullApplicationStatus{"applied", true},
		ID:          int32(1),
	}

	mock.ExpectExec("UPDATE Application SET").WithArgs(param.CoverLetter, param.Resume, param.Status, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateApplication(context.TODO(), param)
	require.NoError(t, err)
}
