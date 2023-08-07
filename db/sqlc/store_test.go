package db

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestExecTx(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	store := NewStore(db)
	mock.ExpectBegin()
	mock.ExpectCommit()

	fn := func(q *Queries) error {
		return nil
	}

	err = store.execTx(ctx, fn)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
