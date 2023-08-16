package db

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomJobSeekerSkillParams() InsertJobSeekerSkillParams {
	return InsertJobSeekerSkillParams{
		JobSeekerID:    int32(gofakeit.Number(1, 1000)),
		TechnicalSkill: Technicalskills(gofakeit.Word()),
		PassiveSkill:   Passiveskills(gofakeit.Word()),
	}
}

func TestDeleteJobSeekerSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	params := DeleteJobSeekerSkillParams{
		JobSeekerID:    1,
		TechnicalSkill: "some_skill",
		PassiveSkill:   "some_passive_skill",
	}
	mock.ExpectExec("DELETE FROM JobSeekerSkill WHERE").WithArgs(params.JobSeekerID, params.TechnicalSkill, params.PassiveSkill).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteJobSeekerSkill(context.TODO(), params)
	require.NoError(t, err)
}

func TestInsertJobSeekerSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomJobSeekerSkillParams()
	mock.ExpectExec("INSERT INTO JobSeekerSkill").WithArgs(param.JobSeekerID, param.TechnicalSkill, param.PassiveSkill).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertJobSeekerSkill(context.TODO(), param)
	require.NoError(t, err)
}

func TestSelectJobSeekerSkillsByJobSeekerID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	jobSeekerID := int32(1)
	rows := sqlmock.NewRows([]string{"job_seeker_id", "technical_skill", "passive_skill"}).
		AddRow(jobSeekerID, "some_skill", "some_passive_skill")

	mock.ExpectQuery("SELECT (.+) FROM JobSeekerSkill WHERE job_seeker_id =").WithArgs(jobSeekerID).WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectJobSeekerSkillsByJobSeekerID(context.TODO(), jobSeekerID)
	require.NoError(t, err)
}
