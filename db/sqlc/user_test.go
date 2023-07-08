package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Name:  sql.NullString{String: "John Doe", Valid: true},
		Email: sql.NullString{String: "johndoe@example.com", Valid: true},
		Phone: sql.NullString{String: "1234567890", Valid: true},
	}

	id, err := testQueries.CreateUser(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteUser(t *testing.T) {
	id := int32(1) // replace with a valid id from your test db
	err := testQueries.DeleteUser(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	users, err := testQueries.GetUsers(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestUpdateUser(t *testing.T) {
	arg := UpdateUserParams{
		Name:  sql.NullString{String: "Jane Doe", Valid: true},
		Email: sql.NullString{String: "janedoe@example.com", Valid: true},
		Phone: sql.NullString{String: "0987654321", Valid: true},
		ID:    1, 
	}

	err := testQueries.UpdateUser(context.TODO(), arg)
	require.NoError(t, err)
}
