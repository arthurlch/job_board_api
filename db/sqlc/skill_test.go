package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomSkill(t *testing.T) sql.NullString {
	return sql.NullString{String: gofakeit.Name(), Valid: true}
}

func TestDeleteSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Skill WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteSkill(context.TODO(), id)
	require.NoError(t, err)
}

func TestInsertSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	name := createRandomSkill(t)
	mock.ExpectExec("INSERT INTO Skill").WithArgs(name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertSkill(context.TODO(), name)
	require.NoError(t, err)
}

func TestSelectAllSkills(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Skill1").
		AddRow(2, "Skill2")

	mock.ExpectQuery("SELECT id, name FROM Skill").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllSkills(context.TODO())
	require.NoError(t, err)
}

func TestSelectSkillByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, "Skill1")

	mock.ExpectQuery("SELECT id, name FROM Skill WHERE id =").WithArgs(id).WillReturnRows(row)

	q := New(db)
	_, err = q.SelectSkillByID(context.TODO(), id)
	require.NoError(t, err)
}

func TestUpdateSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateSkillParams{
		Name: sql.NullString{String: "SkillUpdated", Valid: true},
		ID:   int32(1),
	}

	mock.ExpectExec("UPDATE Skill SET name =").WithArgs(param.Name, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateSkill(context.TODO(), param)
	require.NoError(t, err)
}
