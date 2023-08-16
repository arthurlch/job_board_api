package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomEducationParams(t *testing.T) InsertEducationParams {
	return InsertEducationParams{
		JobSeekerID:   sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		InstitutionID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Degree:        sql.NullString{String: gofakeit.BS(), Valid: true},
		FieldOfStudy:  sql.NullString{String: gofakeit.JobTitle(), Valid: true},
		StartDate:     sql.NullTime{Time: gofakeit.Date(), Valid: true},
		EndDate:       sql.NullTime{Time: gofakeit.Date(), Valid: true},
	}
}

func TestInsertEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomEducationParams(t)
	mock.ExpectExec("INSERT INTO Education").WithArgs(param.JobSeekerID, param.InstitutionID, param.Degree, param.FieldOfStudy, param.StartDate, param.EndDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertEducation(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Education WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteEducation(context.TODO(), id)
	require.NoError(t, err)
}

func TestGetEducations(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "institution_id", "degree", "field_of_study", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(1, 1, 2, "degree", "field_of_study", gofakeit.Date(), gofakeit.Date(), gofakeit.Date(), gofakeit.Date())

	mock.ExpectQuery("SELECT (.+) FROM Education").WillReturnRows(rows)

	q := New(db)
	_, err = q.GetEducations(context.TODO())
	require.NoError(t, err)
}

func TestUpdateEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateEducationParams{
		InstitutionID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Degree:        sql.NullString{String: gofakeit.BS(), Valid: true},
		FieldOfStudy:  sql.NullString{String: gofakeit.JobTitle(), Valid: true},
		StartDate:     sql.NullTime{Time: gofakeit.Date(), Valid: true},
		EndDate:       sql.NullTime{Time: gofakeit.Date(), Valid: true},
		ID:            int32(1),
	}

	mock.ExpectExec("UPDATE Education SET").WithArgs(param.InstitutionID, param.Degree, param.FieldOfStudy, param.StartDate, param.EndDate, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateEducation(context.TODO(), param)
	require.NoError(t, err)
}
