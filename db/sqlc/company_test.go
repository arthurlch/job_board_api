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

func createRandomCompany(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) Company {
	arg := CreateCompanyParams{
		UserID:      sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Name:        sql.NullString{String: gofakeit.Company(), Valid: true},
		Email:       sql.NullString{String: gofakeit.Email(), Valid: true},
		Phone:       sql.NullString{String: gofakeit.Phone(), Valid: true},
		Website:     sql.NullString{String: gofakeit.URL(), Valid: true},
		Logo:        sql.NullString{String: gofakeit.ImageURL(100, 100), Valid: true}, 
		Description: sql.NullString{String: gofakeit.HipsterSentence(10), Valid: true},
	}

	mock.ExpectQuery("INSERT INTO \"Company\"").
		WithArgs(arg.UserID, arg.Name, arg.Email, arg.Phone, arg.Website, arg.Logo, arg.Description).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	companyID, err := q.CreateCompany(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), companyID)

	company := Company{
		ID:          companyID,
		UserID:      arg.UserID,
		Name:        arg.Name,
		Email:       arg.Email,
		Phone:       arg.Phone,
		Website:     arg.Website,
		Logo:        arg.Logo,
		Description: arg.Description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}

	return company
}

func TestCreateCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomCompany(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCompanies(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	company := createRandomCompany(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "email", "phone", "website", "logo", "description", "created_at", "updated_at"}).
		AddRow(company.ID, company.UserID, company.Name, company.Email, company.Phone, company.Website, company.Logo, company.Description, company.CreatedAt, company.UpdatedAt)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	companies, err := q.GetCompanies(context.Background())
	require.NoError(t, err)

	require.Len(t, companies, 1)

	require.Equal(t, company, companies[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	company := createRandomCompany(t, db, mock)

	arg := UpdateCompanyParams{
		Name:        sql.NullString{String: "Updated Company", Valid: true},
		Email:       sql.NullString{String: "updated.company@example.com", Valid: true},
		Phone:       sql.NullString{String: "0987654321", Valid: true},
		Website:     sql.NullString{String: "https://updated.company", Valid: true},
		Logo:        sql.NullString{String: "https://updated.company/logo.png", Valid: true},
		Description: sql.NullString{String: "Updated Description", Valid: true},
		ID:          company.ID,
	}

	mock.ExpectExec("UPDATE \"Company\"").
		WithArgs(arg.Name, arg.Email, arg.Phone, arg.Website, arg.Logo, arg.Description, arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateCompany(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	company := createRandomCompany(t, db, mock)

	mock.ExpectExec("DELETE FROM \"Company\"").
		WithArgs(company.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteCompany(context.TODO(), company.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
