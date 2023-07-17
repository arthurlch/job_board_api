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

func createRandomEducation(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) Education {
	arg := CreateEducationParams{
		JobSeekerID:  sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Institution:  sql.NullString{String: gofakeit.Company(), Valid: true}, 
		Degree:       sql.NullString{String: gofakeit.BS(), Valid: true},  
		FieldOfStudy: sql.NullString{String: gofakeit.BS(), Valid: true},  
		StartDate:    sql.NullTime{Time: gofakeit.Date(), Valid: true},
		EndDate:      sql.NullTime{Time: gofakeit.Date(), Valid: true},
	}


	mock.ExpectQuery("INSERT INTO \"Education\"").
		WithArgs(arg.JobSeekerID, arg.Institution, arg.Degree, arg.FieldOfStudy, arg.StartDate, arg.EndDate).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	q := New(db)
	educationID, err := q.CreateEducation(context.TODO(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(1), educationID)

	education := Education{
		ID:            educationID,
		JobSeekerID:   arg.JobSeekerID,
		Institution:   arg.Institution,
		Degree:        arg.Degree,
		FieldOfStudy:  arg.FieldOfStudy,
		StartDate:     arg.StartDate,
		EndDate:       arg.EndDate,
		CreatedAt:     sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     sql.NullTime{Time: time.Now(), Valid: true},
	}

	return education
}

func TestCreateEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	createRandomEducation(t, db, mock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetEducations(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	education := createRandomEducation(t, db, mock)

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "institution", "degree", "field_of_study", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(education.ID, education.JobSeekerID, education.Institution, education.Degree, education.FieldOfStudy, education.StartDate, education.EndDate, education.CreatedAt, education.UpdatedAt)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	q := New(db)
	educations, err := q.GetEducations(context.Background())
	require.NoError(t, err)

	require.Len(t, educations, 1)

	require.Equal(t, education, educations[0])
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	education := createRandomEducation(t, db, mock)

	arg := UpdateEducationParams{
		Institution:  sql.NullString{String: "Updated University", Valid: true},
		Degree:       sql.NullString{String: "Master", Valid: true},
		FieldOfStudy: sql.NullString{String: "Computer Engineering", Valid: true},
		StartDate:    sql.NullTime{Time: time.Now(), Valid: true},
		EndDate:      sql.NullTime{Time: time.Now(), Valid: true},
		ID:           education.ID,
	}

	mock.ExpectExec("UPDATE \"Education\"").
		WithArgs(arg.Institution, arg.Degree, arg.FieldOfStudy, arg.StartDate, arg.EndDate, arg.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateEducation(context.TODO(), arg)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	education := createRandomEducation(t, db, mock)

	mock.ExpectExec("DELETE FROM \"Education\"").
		WithArgs(education.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	q := New(db)
	err = q.DeleteEducation(context.TODO(), education.ID)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
