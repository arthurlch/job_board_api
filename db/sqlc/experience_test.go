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

func createRandomExperience(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) Experience {
	arg := CreateExperienceParams{
		JobSeekerID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Title:       sql.NullString{String: gofakeit.JobTitle(), Valid: true},
		Company:     sql.NullString{String: gofakeit.Company(), Valid: true},
		Location:    sql.NullString{String: gofakeit.City(), Valid: true},
		StartDate:   sql.NullTime{Time: gofakeit.Date(), Valid: true},
		EndDate:     sql.NullTime{Time: gofakeit.Date(), Valid: true},
		Description: sql.NullString{String: gofakeit.JobDescriptor(), Valid: true},
}

	mock.ExpectQuery("INSERT INTO Experience").
		WithArgs(arg.JobSeekerID, arg.Title, arg.Company, arg.Location, arg.StartDate, arg.EndDate, arg.Description).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	experienceID, err := q.CreateExperience(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), experienceID)

	experience := Experience{
		ID:          experienceID,
		JobSeekerID: arg.JobSeekerID,
		Title:       arg.Title,
		Company:     arg.Company,
		Location:    arg.Location,
		StartDate:   arg.StartDate,
		EndDate:     arg.EndDate,
		Description: arg.Description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}

	return experience
}

func TestCreateExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomExperience(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExperiences(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	experience := createRandomExperience(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "title", "company", "location", "start_date", "end_date", "description", "created_at", "updated_at"}).
		AddRow(experience.ID, experience.JobSeekerID, experience.Title, experience.Company, experience.Location, experience.StartDate, experience.EndDate, experience.Description, experience.CreatedAt, experience.UpdatedAt)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	experiences, err := q.GetExperiences(context.Background())
	require.NoError(t, err)

	require.Len(t, experiences, 1)

	require.Equal(t, experience, experiences[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	experience := createRandomExperience(t, db, mock)

	arg := UpdateExperienceParams{
		Title:       sql.NullString{String: "Senior Software Engineer", Valid: true},
		Company:     sql.NullString{String: "Updated Tech Co.", Valid: true},
		Location:    sql.NullString{String: "Updated Location", Valid: true},
		StartDate:   sql.NullTime{Time: time.Now(), Valid: true},
		EndDate:     sql.NullTime{Time: time.Now(), Valid: true},
		Description: sql.NullString{String: "Updated software", Valid: true},
		ID:          experience.ID,
	}

	mock.ExpectExec("UPDATE Experience").
		WithArgs(arg.Title, arg.Company, arg.Location, arg.StartDate, arg.EndDate, arg.Description, arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateExperience(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	experience := createRandomExperience(t, db, mock)

	mock.ExpectExec("DELETE FROM Experience").
		WithArgs(experience.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteExperience(context.TODO(), experience.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
