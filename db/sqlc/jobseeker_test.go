package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateJobSeeker(t *testing.T) {
	arg := CreateJobSeekerParams{
		UserID: sql.NullInt32{Int32: 1, Valid: true},
		Resume: sql.NullString{String: "resume.pdf", Valid: true},
		Skills: []string{"Go", "SQL", "Python"},
	}

	id, err := testQueries.CreateJobSeeker(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteJobSeeker(t *testing.T) {
	id := int32(1) 
	err := testQueries.DeleteJobSeeker(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetJobSeekers(t *testing.T) {
	jobSeekers, err := testQueries.GetJobSeekers(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, jobSeekers)
}

func TestUpdateJobSeeker(t *testing.T) {
	arg := UpdateJobSeekerParams{
		UserID: sql.NullInt32{Int32: 2, Valid: true}, 
		Resume: sql.NullString{String: "resume_v2.pdf", Valid: true},
		Skills: []string{"Go", "SQL", "Python", "Java"},
		ID:     1, 
	}

	err := testQueries.UpdateJobSeeker(context.TODO(), arg)
	require.NoError(t, err)
}
