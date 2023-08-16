package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomJobViewParams(t *testing.T) InsertJobViewParams {
	return InsertJobViewParams{
		JobID:     int32(gofakeit.Number(1, 1000)),
		ViewCount: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
	}
}

func TestInsertJobView(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomJobViewParams(t)
	mock.ExpectExec("INSERT INTO JobViews").WithArgs(param.JobID, param.ViewCount).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertJobView(context.TODO(), param)
	require.NoError(t, err)
}

func TestSelectJobViewsByJobID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	jobID := int32(1)
	row := sqlmock.NewRows([]string{"job_id", "view_count"}).
		AddRow(jobID, 100)

	mock.ExpectQuery("SELECT job_id, view_count FROM JobViews WHERE job_id =").WithArgs(jobID).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectJobViewsByJobID(context.TODO(), jobID)
	require.NoError(t, err)
}

func TestUpdateJobViews(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateJobViewsParams{
		ViewCount: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		JobID:     int32(1),
	}

	mock.ExpectExec("UPDATE JobViews SET view_count =").WithArgs(param.ViewCount, param.JobID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJobViews(context.TODO(), param)
	require.NoError(t, err)
}
