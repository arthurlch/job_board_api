package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomInstitution(t *testing.T) sql.NullString {
	return sql.NullString{String: gofakeit.Company(), Valid: true}
}

func TestDeleteInstitution(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Institution WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteInstitution(context.TODO(), id)
	require.NoError(t, err)
}

func TestInsertInstitution(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	name := createRandomInstitution(t)
	mock.ExpectExec("INSERT INTO Institution").WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertInstitution(context.TODO(), name)
	require.NoError(t, err)
}

func TestSelectAllInstitutions(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Institution1")

	mock.ExpectQuery("SELECT (.+) FROM Institution").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllInstitutions(context.TODO())
	require.NoError(t, err)
}

func TestSelectInstitutionByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Institution1")

	mock.ExpectQuery("SELECT (.+) FROM Institution WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectInstitutionByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateInstitution(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateInstitutionParams{
		Name: sql.NullString{String: "NewInstitution", Valid: true},
		ID:   int32(1),
	}

	mock.ExpectExec("UPDATE Institution SET name =").WithArgs(param.Name, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateInstitution(context.TODO(), param)
	require.NoError(t, err)
}
