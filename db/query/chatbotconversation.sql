-- Insert (Create)
-- name: InsertChatbotConversation :exec
INSERT INTO chatbotconversation (chatbot_interview_id, sender_type, content)
VALUES ($1, $2, $3)
RETURNING id;

-- Select All (Read)
-- name: SelectAllChatbotConversations :many
SELECT * FROM chatbotconversation;

-- Select by chatbot_interview_id (Read)
-- name: SelectChatbotConversationsByInterviewID :many
SELECT * FROM chatbotconversation WHERE chatbot_interview_id = $1;

-- Update (Update)
-- name: UpdateChatbotConversation :exec
UPDATE chatbotconversation SET chatbot_interview_id = $1, sender_type = $2, content = $3 WHERE id = $4;

-- Delete (Delete)
-- name: DeleteChatbotConversation :exec
DELETE FROM chatbotconversation WHERE id = $1;