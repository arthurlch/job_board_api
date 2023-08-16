package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomExperienceType() string {
	return gofakeit.Word()
}

func TestDeleteExperienceType(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM ExperienceType WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteExperienceType(context.TODO(), id)
	require.NoError(t, err)
}

func TestInsertExperienceType(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	name := sql.NullString{String: createRandomExperienceType(), Valid: true}
	mock.ExpectExec("INSERT INTO ExperienceType").WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertExperienceType(context.TODO(), name)
	require.NoError(t, err)
}

func TestSelectAllExperienceTypes(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, createRandomExperienceType())

	mock.ExpectQuery("SELECT id, name FROM ExperienceType").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllExperienceTypes(context.TODO())
	require.NoError(t, err)
}

func TestSelectExperienceTypeByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, createRandomExperienceType())

	mock.ExpectQuery("SELECT id, name FROM ExperienceType WHERE id =").WithArgs(id).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectExperienceTypeByID(context.TODO(), id)
	require.NoError(t, err)
}

func TestUpdateExperienceType(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateExperienceTypeParams{
		Name: sql.NullString{String: createRandomExperienceType(), Valid: true},
		ID:   int32(1),
	}

	mock.ExpectExec("UPDATE ExperienceType SET name =").WithArgs(param.Name, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateExperienceType(context.TODO(), param)
	require.NoError(t, err)
}
