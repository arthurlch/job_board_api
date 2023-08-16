package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomJobSeekerParams(t *testing.T) CreateJobSeekerParams {
	return CreateJobSeekerParams{
		UserID:  sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
	}
}

func TestCreateJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomJobSeekerParams(t)
	mock.ExpectExec("INSERT INTO JobSeeker").WithArgs(param.UserID, pq.Array(param.Column2)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	_, err = q.CreateJobSeeker(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM JobSeeker WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteJobSeeker(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetJobSeekers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "skills", "created_at", "updated_at"}).
		AddRow(1, 1, pq.Array([]string{"skill1", "skill2"}), time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM JobSeeker").WillReturnRows(rows)

	q := New(db)
	_, err = q.GetJobSeekers(context.TODO())
	require.NoError(t, err)
}

func TestUpdateJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateJobSeekerParams{
		UserID:  sql.NullInt32{Int32: int32(1), Valid: true},
		ID:      int32(1),
	}

	mock.ExpectExec("UPDATE JobSeeker SET").WithArgs(param.UserID, pq.Array(param.Column2), param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJobSeeker(context.TODO(), param)
	require.NoError(t, err)
}
