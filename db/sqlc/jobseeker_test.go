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

func createRandomJobSeeker(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) Jobseeker {
	arg := CreateJobSeekerParams{
		UserID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Resume: sql.NullString{String: gofakeit.Word() + ".pdf", Valid: true},
		Skills: []string{"Go", "SQL", "Python"},
	}

	mock.ExpectQuery("INSERT INTO JobSeeker").
		WithArgs(arg.UserID, arg.Resume, pq.Array(arg.Skills)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	jobSeekerID, err := q.CreateJobSeeker(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), jobSeekerID)

	jobSeeker := Jobseeker{
		ID:        jobSeekerID,
		UserID:    arg.UserID,
		Resume:    arg.Resume,
		Skills:    arg.Skills,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return jobSeeker
}

func TestCreateJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomJobSeeker(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetJobSeekers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	jobSeeker := createRandomJobSeeker(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "user_id", "resume", "skills", "created_at", "updated_at"}).
		AddRow(jobSeeker.ID, jobSeeker.UserID, jobSeeker.Resume, pq.Array(jobSeeker.Skills), jobSeeker.CreatedAt, jobSeeker.UpdatedAt)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	jobSeekers, err := q.GetJobSeekers(context.Background())
	require.NoError(t, err)

	require.Len(t, jobSeekers, 1)

	require.Equal(t, jobSeeker, jobSeekers[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	jobSeeker := createRandomJobSeeker(t, db, mock)

	arg := UpdateJobSeekerParams{
		UserID: sql.NullInt32{Int32: 2, Valid: true},
		Resume: sql.NullString{String: "resume_v2.pdf", Valid: true},
		Skills: []string{"Go", "SQL", "Python", "Java"},
		ID:     jobSeeker.ID,
	}

	mock.ExpectExec("UPDATE JobSeeker").
		WithArgs(arg.UserID, arg.Resume, pq.Array(arg.Skills), arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJobSeeker(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteJobSeeker(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	jobSeeker := createRandomJobSeeker(t, db, mock)

	mock.ExpectExec("DELETE FROM JobSeeker").
		WithArgs(jobSeeker.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteJobSeeker(context.TODO(), jobSeeker.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
