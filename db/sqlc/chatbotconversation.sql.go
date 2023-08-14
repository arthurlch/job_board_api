// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: chatbotconversation.sql

package db

import (
	"context"
	"database/sql"
)

const deleteChatbotConversation = `-- name: DeleteChatbotConversation :exec
DELETE FROM chatbotconversation WHERE id = $1
`

// Delete (Delete)
func (q *Queries) DeleteChatbotConversation(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteChatbotConversationStmt, deleteChatbotConversation, id)
	return err
}

const insertChatbotConversation = `-- name: InsertChatbotConversation :exec
INSERT INTO chatbotconversation (chatbot_interview_id, sender_type, content)
VALUES ($1, $2, $3)
RETURNING id
`

type InsertChatbotConversationParams struct {
	ChatbotInterviewID sql.NullInt32  `json:"chatbot_interview_id"`
	SenderType         sql.NullString `json:"sender_type"`
	Content            string         `json:"content"`
}

// Insert (Create)
func (q *Queries) InsertChatbotConversation(ctx context.Context, arg InsertChatbotConversationParams) error {
	_, err := q.exec(ctx, q.insertChatbotConversationStmt, insertChatbotConversation, arg.ChatbotInterviewID, arg.SenderType, arg.Content)
	return err
}

const selectAllChatbotConversations = `-- name: SelectAllChatbotConversations :many
SELECT id, chatbot_interview_id, sender_type, content, created_at FROM chatbotconversation
`

// Select All (Read)
func (q *Queries) SelectAllChatbotConversations(ctx context.Context) ([]Chatbotconversation, error) {
	rows, err := q.query(ctx, q.selectAllChatbotConversationsStmt, selectAllChatbotConversations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chatbotconversation
	for rows.Next() {
		var i Chatbotconversation
		if err := rows.Scan(
			&i.ID,
			&i.ChatbotInterviewID,
			&i.SenderType,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectChatbotConversationsByInterviewID = `-- name: SelectChatbotConversationsByInterviewID :many
SELECT id, chatbot_interview_id, sender_type, content, created_at FROM chatbotconversation WHERE chatbot_interview_id = $1
`

// Select by chatbot_interview_id (Read)
func (q *Queries) SelectChatbotConversationsByInterviewID(ctx context.Context, chatbotInterviewID sql.NullInt32) ([]Chatbotconversation, error) {
	rows, err := q.query(ctx, q.selectChatbotConversationsByInterviewIDStmt, selectChatbotConversationsByInterviewID, chatbotInterviewID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chatbotconversation
	for rows.Next() {
		var i Chatbotconversation
		if err := rows.Scan(
			&i.ID,
			&i.ChatbotInterviewID,
			&i.SenderType,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChatbotConversation = `-- name: UpdateChatbotConversation :exec
UPDATE chatbotconversation SET chatbot_interview_id = $1, sender_type = $2, content = $3 WHERE id = $4
`

type UpdateChatbotConversationParams struct {
	ChatbotInterviewID sql.NullInt32  `json:"chatbot_interview_id"`
	SenderType         sql.NullString `json:"sender_type"`
	Content            string         `json:"content"`
	ID                 int32          `json:"id"`
}

// Update (Update)
func (q *Queries) UpdateChatbotConversation(ctx context.Context, arg UpdateChatbotConversationParams) error {
	_, err := q.exec(ctx, q.updateChatbotConversationStmt, updateChatbotConversation,
		arg.ChatbotInterviewID,
		arg.SenderType,
		arg.Content,
		arg.ID,
	)
	return err
}
