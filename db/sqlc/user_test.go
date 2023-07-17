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


func createRandomUser(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) User {
	arg := CreateUserParams{
		Name:  sql.NullString{String: gofakeit.Name(), Valid: true},
		Email: sql.NullString{String: gofakeit.Email(), Valid: true},
		Phone: sql.NullString{String: gofakeit.Phone(), Valid: true},
	}

	mock.ExpectQuery("INSERT INTO \"User\"").
		WithArgs(arg.Name, arg.Email, arg.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	userID, err := q.CreateUser(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), userID)

	user := User{
		ID:        userID,
		Name:      arg.Name,
		Email:     arg.Email,
		Phone:     arg.Phone,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		}

	return user
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomUser(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	user := createRandomUser(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	users, err := q.GetUsers(context.Background())
	require.NoError(t, err)

	require.Len(t, users, 1)

	require.Equal(t, user, users[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	user := createRandomUser(t, db, mock)

	arg := UpdateUserParams{
		Name:  sql.NullString{String: "Updated User", Valid: true},
		Email: sql.NullString{String: "updated.user@example.com", Valid: true},
		Phone: sql.NullString{String: "0987654321", Valid: true},
		ID:    user.ID,
	}

	mock.ExpectExec("UPDATE \"User\"").
		WithArgs(arg.Name, arg.Email, arg.Phone, arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateUser(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	user := createRandomUser(t, db, mock)

	mock.ExpectExec("DELETE FROM \"User\"").
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteUser(context.TODO(), user.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
