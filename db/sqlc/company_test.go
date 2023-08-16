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

func createRandomCompanyParams(t *testing.T) InsertCompanyParams {
	return InsertCompanyParams{
		UserID:      sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Name:        gofakeit.Company(),
		Email:       gofakeit.Email(),
		Phone:       sql.NullString{String: gofakeit.Phone(), Valid: true},
		Website:     sql.NullString{String: gofakeit.URL(), Valid: true},
		Logo:        sql.NullString{String: gofakeit.URL(), Valid: true},
		Description: sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
	}
}

func TestDeleteCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Company WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteCompany(context.TODO(), id)
	require.NoError(t, err)
}

func TestInsertCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomCompanyParams(t)
	mock.ExpectExec("INSERT INTO Company").WithArgs(param.UserID, param.Name, param.Email, param.Phone, param.Website, param.Logo, param.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertCompany(context.TODO(), param)
	require.NoError(t, err)
}

func TestSelectAllCompanies(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "email", "phone", "website", "logo", "description", "created_at", "updated_at"}).
		AddRow(1, 1, "name", "email", "phone", "website", "logo", "description", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Company").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllCompanies(context.TODO())
	require.NoError(t, err)
}

func TestSelectCompanyByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "user_id", "name", "email", "phone", "website", "logo", "description", "created_at", "updated_at"}).
		AddRow(1, 1, "name", "email", "phone", "website", "logo", "description", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Company WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectCompanyByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateCompanyParams{
		Name:        gofakeit.Company(),
		Email:       gofakeit.Email(),
		Phone:       sql.NullString{String: gofakeit.Phone(), Valid: true},
		Website:     sql.NullString{String: gofakeit.URL(), Valid: true},
		Logo:        sql.NullString{String: gofakeit.URL(), Valid: true},
		Description: sql.NullString{String: gofakeit.LoremIpsumParagraph(3, 3, 3, " "), Valid: true},
		ID:          int32(1),
	}

	mock.ExpectExec("UPDATE Company SET").WithArgs(param.Name, param.Email, param.Phone, param.Website, param.Logo, param.Description, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateCompany(context.TODO(), param)
	require.NoError(t, err)
}
