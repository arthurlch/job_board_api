package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomJob(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) Job {
	arg := CreateJobParams{
		Title:        sql.NullString{String: gofakeit.JobTitle(), Valid: true},
		Description:  sql.NullString{String: gofakeit.JobDescriptor(), Valid: true},
		Requirements: sql.NullString{String: "Knowledge in " + gofakeit.ProgrammingLanguage(), Valid: true},
		Location:     sql.NullString{String: gofakeit.City(), Valid: true},
		Salary: sql.NullInt32{Int32: int32(gofakeit.Number(60000,150000)), Valid: true},
		CompanyID:    sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
	}

	mock.ExpectQuery("INSERT INTO \"Job\"").
		WithArgs(arg.Title, arg.Description, arg.Requirements, arg.Location, arg.Salary, arg.CompanyID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	jobID, err := q.CreateJob(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), jobID)

	job := Job{
		ID:           jobID,
		Title:        arg.Title,
		Description:  arg.Description,
		Requirements: arg.Requirements,
		Location:     arg.Location,
		Salary:       arg.Salary,
		CompanyID:    arg.CompanyID,
	}

	return job
}

func TestCreateJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomJob(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetJobs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	job := createRandomJob(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "requirements", "location", "salary", "company_id"}).
		AddRow(job.ID, job.Title, job.Description, job.Requirements, job.Location, job.Salary, job.CompanyID)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	jobs, err := q.GetJobs(context.Background())
	require.NoError(t, err)

	require.Len(t, jobs, 1)

	require.Equal(t, job, jobs[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	job := createRandomJob(t, db, mock)

	arg := UpdateJobParams{
		Title:        sql.NullString{String: "Senior Software Engineer", Valid: true},
		Description:  sql.NullString{String: "Develop and maintain software applications", Valid: true},
		Requirements: sql.NullString{String: "Knowledge in Go and Python", Valid: true},
		Location:     sql.NullString{String: "Remote", Valid: true},
		Salary: sql.NullInt32{Int32: int32(gofakeit.Number(60000,150000)), Valid: true},
		ID:           job.ID,
	}

	mock.ExpectExec("UPDATE \"Job\"").
		WithArgs(arg.Title, arg.Description, arg.Requirements, arg.Location, arg.Salary, arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJob(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	job := createRandomJob(t, db, mock)

	mock.ExpectExec("DELETE FROM \"Job\"").
		WithArgs(job.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteJob(context.TODO(), job.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
