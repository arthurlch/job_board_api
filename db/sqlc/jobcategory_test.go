package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomJobCategory(t *testing.T) sql.NullString {
	return sql.NullString{String: gofakeit.Word(), Valid: true}
}

func TestInsertJobCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	name := createRandomJobCategory(t)
	mock.ExpectExec("INSERT INTO JobCategory").WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertJobCategory(context.TODO(), name)
	require.NoError(t, err)
}

func TestDeleteJobCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM JobCategory WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteJobCategory(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllJobCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Software Engineer")

	mock.ExpectQuery("SELECT (.+) FROM JobCategory").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllJobCategories(context.TODO())
	require.NoError(t, err)
}

func TestSelectJobCategoryByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Software Engineer")

	mock.ExpectQuery("SELECT (.+) FROM JobCategory WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectJobCategoryByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateJobCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateJobCategoryParams{
		Name: sql.NullString{String: "Software Developer", Valid: true},
		ID:   int32(1),
	}

	mock.ExpectExec("UPDATE JobCategory SET").WithArgs(param.Name, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateJobCategory(context.TODO(), param)
	require.NoError(t, err)
}
