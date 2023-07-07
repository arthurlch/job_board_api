// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one

INSERT INTO "User" (
  name,
  email,
  phone,
  created_at,
  updated_at
) VALUES (
  $1, -- name
  $2, -- email
  $3, -- phone
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id
`

type CreateUserParams struct {
	Name  sql.NullString `json:"name"`
	Email sql.NullString `json:"email"`
	Phone sql.NullString `json:"phone"`
}

// user.sql
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Email, arg.Phone)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec

DELETE FROM "User"
WHERE id = $1
`

// Specify the condition for updating, such as the "id" column
func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, email, phone, created_at, updated_at FROM "User"
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE "User"
SET name = $1, email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4
`

type UpdateUserParams struct {
	Name  sql.NullString `json:"name"`
	Email sql.NullString `json:"email"`
	Phone sql.NullString `json:"phone"`
	ID    int32          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.ID,
	)
	return err
}
