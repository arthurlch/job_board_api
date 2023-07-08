package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateExperience(t *testing.T) {
	arg := CreateExperienceParams{
		JobSeekerID: sql.NullInt32{Int32: 1, Valid: true},
		Title:       sql.NullString{String: "Software Engineer", Valid: true},
		Company:     sql.NullString{String: "Tech Corp", Valid: true},
		Location:    sql.NullString{String: "New York, NY", Valid: true},
		StartDate:   sql.NullTime{Time: time.Now().AddDate(-2, 0, 0), Valid: true},
		EndDate:     sql.NullTime{Time: time.Now(), Valid: true},
		Description: sql.NullString{String: "Developing web applications", Valid: true},
	}

	id, err := testQueries.CreateExperience(context.TODO(), arg)
	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestDeleteExperience(t *testing.T) {
	id := int32(1) 
	err := testQueries.DeleteExperience(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetExperiences(t *testing.T) {
	experiences, err := testQueries.GetExperiences(context.TODO())
	require.NoError(t, err)
	assert.NotEmpty(t, experiences)
}

func TestUpdateExperience(t *testing.T) {
	arg := UpdateExperienceParams{
		Title:       sql.NullString{String: "Senior Software Engineer", Valid: true},
		Company:     sql.NullString{String: "Updated Tech Corp", Valid: true},
		Location:    sql.NullString{String: "San Francisco, CA", Valid: true},
		StartDate:   sql.NullTime{Time: time.Now().AddDate(-3, 0, 0), Valid: true},
		EndDate:     sql.NullTime{Time: time.Now().AddDate(-1, 0, 0), Valid: true},
		Description: sql.NullString{String: "Developing advanced web applications", Valid: true},
		ID:          1, 
	}

	err := testQueries.UpdateExperience(context.TODO(), arg)
	require.NoError(t, err)
}
