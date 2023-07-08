package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEducation(t *testing.T) {
	arg := CreateEducationParams{
		JobSeekerID:  sql.NullInt32{Int32: 1, Valid: true},
		Institution:  sql.NullString{String: "Test University", Valid: true},
		Degree:       sql.NullString{String: "B.S. in Computer Science", Valid: true},
		FieldOfStudy: sql.NullString{String: "Computer Science", Valid: true},
		StartDate:    sql.NullTime{Time: time.Now().AddDate(-4, 0, 0), Valid: true},
		EndDate:      sql.NullTime{Time: time.Now(), Valid: true},
	}

	id, err := testQueries.CreateEducation(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteEducation(t *testing.T) {
	id := int32(1) 
	err := testQueries.DeleteEducation(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetEducations(t *testing.T) {
	educations, err := testQueries.GetEducations(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, educations)
}

func TestUpdateEducation(t *testing.T) {
	arg := UpdateEducationParams{
		Institution:  sql.NullString{String: "Updated University", Valid: true},
		Degree:       sql.NullString{String: "Updated Degree", Valid: true},
		FieldOfStudy: sql.NullString{String: "Updated Field of Study", Valid: true},
		StartDate:    sql.NullTime{Time: time.Now().AddDate(-5, 0, 0), Valid: true},
		EndDate:      sql.NullTime{Time: time.Now().AddDate(-1, 0, 0), Valid: true},
		ID:           1, 
	}

	err := testQueries.UpdateEducation(context.TODO(), arg)
	require.NoError(t, err)
}
