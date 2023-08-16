package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomMessageParams(t *testing.T) InsertMessageParams {
	return InsertMessageParams{
		SenderID:   int32(gofakeit.Number(1, 1000)),
		ReceiverID: int32(gofakeit.Number(1, 1000)),
		Content:    gofakeit.LoremIpsumSentence(10),
		SenderType: sql.NullString{String: "type", Valid: true},
	}
}

func TestDeleteMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM Messages WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteMessage(context.TODO(), id)
	require.NoError(t, err)
}

func TestInsertMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomMessageParams(t)
	mock.ExpectExec("INSERT INTO Messages").WithArgs(param.SenderID, param.ReceiverID, param.Content, param.SenderType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertMessage(context.TODO(), param)
	require.NoError(t, err)
}

func TestSelectMessagesBySenderAndReceiver(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := SelectMessagesBySenderAndReceiverParams{
		SenderID:   int32(1),
		ReceiverID: int32(2),
	}
	rows := sqlmock.NewRows([]string{"id", "sender_id", "receiver_id", "content", "sender_type", "created_at"}).
		AddRow(1, 1, 2, "content", "type", "2023-08-14")

	mock.ExpectQuery("SELECT (.+) FROM Messages WHERE sender_id =").WithArgs(param.SenderID, param.ReceiverID).WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectMessagesBySenderAndReceiver(context.TODO(), param)
	require.NoError(t, err)
}
