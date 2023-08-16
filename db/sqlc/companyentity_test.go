package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomCompanyEntity() (string, int32) {
	return gofakeit.Company(), int32(gofakeit.Number(1, 1000))
}

func TestInsertCompanyEntity(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	name, _ := createRandomCompanyEntity()
	mock.ExpectExec("INSERT INTO CompanyEntity").WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertCompanyEntity(context.TODO(), sql.NullString{String: name, Valid: true})
	require.NoError(t, err)
}

func TestDeleteCompanyEntity(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM CompanyEntity WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteCompanyEntity(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllCompanyEntities(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Company Name")

	mock.ExpectQuery("SELECT (.+) FROM CompanyEntity").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllCompanyEntities(context.TODO())
	require.NoError(t, err)
}

func TestSelectCompanyEntityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Company Name")

	mock.ExpectQuery("SELECT (.+) FROM CompanyEntity WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectCompanyEntityByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateCompanyEntity(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateCompanyEntityParams{
		Name: sql.NullString{String: "New Name", Valid: true},
		ID:   int32(1),
	}

	mock.ExpectExec("UPDATE CompanyEntity SET").WithArgs(param.Name, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateCompanyEntity(context.TODO(), param)
	require.NoError(t, err)
}
