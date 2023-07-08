package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateJob(t *testing.T) {
	arg := CreateJobParams{
		Title:        sql.NullString{String: "Software Engineer", Valid: true},
		Description:  sql.NullString{String: "Develop web applications", Valid: true},
		Requirements: sql.NullString{String: "Experience with Go and SQL", Valid: true},
		Location:     sql.NullString{String: "New York, NY", Valid: true},
		Salary:       sql.NullString{String: "100,000 - 120,000 USD per year", Valid: true},
		CompanyID:    sql.NullInt32{Int32: 1, Valid: true}, 
	}

	id, err := testQueries.CreateJob(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteJob(t *testing.T) {
	id := int32(1) 
	err := testQueries.DeleteJob(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetJobs(t *testing.T) {
	jobs, err := testQueries.GetJobs(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, jobs)
}

func TestUpdateJob(t *testing.T) {
	arg := UpdateJobParams{
		Title:        sql.NullString{String: "Senior Software Engineer", Valid: true},
		Description:  sql.NullString{String: "Develop advanced web applications", Valid: true},
		Requirements: sql.NullString{String: "Experience with Go, SQL, and microservices", Valid: true},
		Location:     sql.NullString{String: "San Francisco, CA", Valid: true},
		Salary:       sql.NullString{String: "120,000 - 140,000 USD per year", Valid: true},
		ID:           1,
	}

	err := testQueries.UpdateJob(context.TODO(), arg)
	require.NoError(t, err)
}
