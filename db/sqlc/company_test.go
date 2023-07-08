package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCompany(t *testing.T) {
	arg := CreateCompanyParams{
		UserID:      sql.NullInt32{Int32: 1, Valid: true},
		Name:        sql.NullString{String: "Test Company", Valid: true},
		Email:       sql.NullString{String: "test@example.com", Valid: true},
		Phone:       sql.NullString{String: "1234567890", Valid: true},
		Website:     sql.NullString{String: "https://example.com", Valid: true},
		Logo:        sql.NullString{String: "logo.png", Valid: true},
		Description: sql.NullString{String: "Test company description", Valid: true},
	}

	id, err := testQueries.CreateCompany(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteCompany(t *testing.T) {
	id := int32(1) 
	err := testQueries.DeleteCompany(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetCompanies(t *testing.T) {
	companies, err := testQueries.GetCompanies(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, companies)
}

func TestUpdateCompany(t *testing.T) {
	arg := UpdateCompanyParams{
		Name:        sql.NullString{String: "Updated Test Company", Valid: true},
		Email:       sql.NullString{String: "updatedtest@example.com", Valid: true},
		Phone:       sql.NullString{String: "0987654321", Valid: true},
		Website:     sql.NullString{String: "https://updatedexample.com", Valid: true},
		Logo:        sql.NullString{String: "updatedlogo.png", Valid: true},
		Description: sql.NullString{String: "Updated test company description", Valid: true},
		ID:          1, 
	}

	err := testQueries.UpdateCompany(context.TODO(), arg)
	require.NoError(t, err)
}
