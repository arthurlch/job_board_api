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

func createRandomUserParams(t *testing.T) InsertUserParams {
	return InsertUserParams{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Phone: sql.NullString{String: gofakeit.Phone(), Valid: true},
		Role:  sql.NullString{String: gofakeit.JobTitle(), Valid: true},
	}
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomUserParams(t)
	mock.ExpectExec("INSERT INTO \"User\"").WithArgs(param.Name, param.Email, param.Phone, param.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertUser(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM \"User\" WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteUser(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "role", "created_at", "updated_at"}).
		AddRow(1, "John Doe", "johndoe@example.com", "1234567890", "Manager", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM \"User\"").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectUsers(context.TODO())
	require.NoError(t, err)
}

func TestSelectUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name", "email", "phone", "role", "created_at", "updated_at"}).
		AddRow(1, "John Doe", "johndoe@example.com", "1234567890", "Manager", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM \"User\" WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectUserByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateUserParams{
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
		Phone: sql.NullString{String: "9876543210", Valid: true},
		Role:  sql.NullString{String: "Developer", Valid: true},
		ID:    int32(1),
	}

	mock.ExpectExec("UPDATE \"User\" SET").WithArgs(param.Name, param.Email, param.Phone, param.Role, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateUser(context.TODO(), param)
	require.NoError(t, err)
}
