package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomJobParams(t *testing.T) CreateJobParams {
	return CreateJobParams{
		Title:        gofakeit.BS(),
		Description:  sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Requirements: sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Location:     sql.NullString{String: gofakeit.City(), Valid: true},
		Salary:       sql.NullInt32{Int32: int32(gofakeit.Number(10000, 100000)), Valid: true},
		CompanyID:    sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
	}
}

func TestCreateJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomJobParams(t)
	mock.ExpectQuery("INSERT INTO Job").WithArgs(param.Title, param.Description, param.Requirements, param.Location, param.Salary, param.CompanyID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	_, err = q.CreateJob(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Job WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteJob(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetJobs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "requirements", "location", "salary", "company_id", "created_at", "updated_at"}).
		AddRow(1, "title", "description", "requirements", "location", 50000, 1, gofakeit.Date(), gofakeit.Date())

	mock.ExpectQuery("SELECT (.+) FROM Job").WillReturnRows(rows)

	q := New(db)
	_, err = q.GetJobs(context.TODO())
	require.NoError(t, err)
}

func TestUpdateJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateJobParams{
		Title:        gofakeit.BS(),
		Description:  sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Requirements: sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		Location:     sql.NullString{String: gofakeit.City(), Valid: true},
		Salary:       sql.NullInt32{Int32: int32(gofakeit.Number(10000, 100000)), Valid: true},
		ID:           int32(1),
	}

	mock.ExpectExec("UPDATE Job SET").WithArgs(param.Title, param.Description, param.Requirements, param.Location, param.Salary, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJob(context.TODO(), param)
	require.NoError(t, err)
}
