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

func createRandomChatbotConversationParams(t *testing.T) InsertChatbotConversationParams {
	return InsertChatbotConversationParams{
		ChatbotInterviewID: sql.NullInt32{Int32: int32(gofakeit.Number(1, 1000)), Valid: true},
		SenderType:         sql.NullString{String: gofakeit.RandomString([]string{"user", "bot"}), Valid: true},
		Content:            gofakeit.LoremIpsumParagraph(3, 3, 3, " "),
	}
}

func TestInsertChatbotConversation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := createRandomChatbotConversationParams(t)
	mock.ExpectExec("INSERT INTO chatbotconversation").WithArgs(param.ChatbotInterviewID, param.SenderType, param.Content).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.InsertChatbotConversation(context.TODO(), param)
	require.NoError(t, err)
}

func TestDeleteChatbotConversation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	id := int32(1)
	mock.ExpectExec("DELETE FROM chatbotconversation WHERE id =").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.DeleteChatbotConversation(context.TODO(), id)
	require.NoError(t, err)
}

func TestSelectAllChatbotConversations(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "chatbot_interview_id", "sender_type", "content", "created_at"}).
		AddRow(1, 1, "user", "content", time.Now())

	mock.ExpectQuery("SELECT (.+) FROM chatbotconversation").WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectAllChatbotConversations(context.TODO())
	require.NoError(t, err)
}

func TestSelectChatbotConversationsByInterviewID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	interviewID := sql.NullInt32{Int32: int32(1), Valid: true}
	rows := sqlmock.NewRows([]string{"id", "chatbot_interview_id", "sender_type", "content", "created_at"}).
		AddRow(1, 1, "user", "content", time.Now())

	mock.ExpectQuery("SELECT (.+) FROM chatbotconversation WHERE chatbot_interview_id =").WithArgs(interviewID).WillReturnRows(rows)

	q := New(db)
	_, err = q.SelectChatbotConversationsByInterviewID(context.TODO(), interviewID)
	require.NoError(t, err)
}

func TestUpdateChatbotConversation(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	param := UpdateChatbotConversationParams{
		ChatbotInterviewID: sql.NullInt32{Int32: int32(1), Valid: true},
		SenderType:         sql.NullString{String: "bot", Valid: true},
		Content:            "content",
		ID:                 int32(1),
	}

	mock.ExpectExec("UPDATE chatbotconversation SET").WithArgs(param.ChatbotInterviewID, param.SenderType, param.Content, param.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	q := New(db)
	err = q.UpdateChatbotConversation(context.TODO(), param)
	require.NoError(t, err)
}
