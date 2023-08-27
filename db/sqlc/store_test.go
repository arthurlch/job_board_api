package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestExecTx(t *testing.T) {
	// Initialize mock DB and sqlmock
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	// Initialize the store with mock DB
	store := NewStore(db)

	// Prepare to mock SQL queries
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Application").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Prepare test data
	testData := InsertApplicationParams{
		JobSeekerID: sql.NullInt32{Int32: 1, Valid: true},
		JobID:       sql.NullInt32{Int32: 2, Valid: true},
		CoverLetter: sql.NullString{String: "Hello", Valid: true},
		Resume:      sql.NullString{String: "Resume.pdf", Valid: true},
		Status:      NullApplicationStatus{ApplicationStatus: "PENDING", Valid: true},
	}

	// Test execTx with InsertApplication
	err = store.execTx(context.Background(), func(q *Queries) error {
		return q.InsertApplication(context.Background(), testData)
	})

	// Assertions
	require.NoError(t, err) // Expect no errors
	require.NoError(t, mock.ExpectationsWereMet()) // All SQL queries were properly mocked
}

