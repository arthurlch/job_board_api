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

func createRandomExperienceParams(t *testing.T) InsertExperienceParams {
	return InsertExperienceParams{
		JobSeekerID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Title:       sql.NullString{String: gofakeit.Sentence(5), Valid: true},
		CompanyID:   sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		Location:    sql.NullString{String: gofakeit.City(), Valid: true},
		StartDate:   sql.NullTime{Time: gofakeit.Date(), Valid: true},
		EndDate:     sql.NullTime{Time: gofakeit.Date(), Valid: true},
		TypeID:      sql.NullInt32{Int32: int32(gofakeit.Number(1, 5)), Valid: true},
		Description: sql.NullString{String: gofakeit.Paragraph(3, 3, 3, " "), Valid: true},
	}
}

func TestInsertExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomExperienceParams(t)
	mock.ExpectExec("INSERT INTO Experience").WithArgs(
		param.JobSeekerID, param.Title, param.CompanyID,
		param.Location, param.StartDate, param.EndDate,
		param.TypeID, param.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertExperience(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Experience WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteExperience(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllExperiences(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "job_seeker_id", "title", "company_id", "location", "start_date", "end_date", "type_id", "description", "created_at", "updated_at"}).
		AddRow(1, 1, "title", 2, "location", time.Now(), time.Now(), 1, "description", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Experience").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllExperiences(context.TODO())
	require.NoError(t, err)
}

func TestSelectExperienceByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "job_seeker_id", "title", "company_id", "location", "start_date", "end_date", "type_id", "description", "created_at", "updated_at"}).
		AddRow(1, 1, "title", 2, "location", time.Now(), time.Now(), 1, "description", time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.+) FROM Experience WHERE id =").WithArgs(1).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectExperienceByID(context.TODO(), 1)
	require.NoError(t, err)
}

func TestUpdateExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateExperienceParams{
		Title:      sql.NullString{String: "title", Valid: true},
		Location:   sql.NullString{String: "location", Valid: true},
		StartDate:  sql.NullTime{Time: time.Now(), Valid: true},
		EndDate:    sql.NullTime{Time: time.Now(), Valid: true},
		TypeID:     sql.NullInt32{Int32: 1, Valid: true},
		Description: sql.NullString{String: "description", Valid: true},
		ID:         int32(1),
	}

	mock.ExpectExec("UPDATE Experience SET").WithArgs(param.Title, param.Location, param.StartDate, param.EndDate, param.TypeID, param.Description, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateExperience(context.TODO(), param)
	require.NoError(t, err)
}
