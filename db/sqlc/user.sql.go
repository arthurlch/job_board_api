// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "User" WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO "User" (name, email, phone, role) VALUES ($1, $2, $3, $4) RETURNING id
`

type InsertUserParams struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Phone sql.NullString `json:"phone"`
	Role  sql.NullString `json:"role"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.exec(ctx, q.insertUserStmt, insertUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Role,
	)
	return err
}

const selectUserByID = `-- name: SelectUserByID :one
SELECT id, name, email, phone, role, created_at, updated_at FROM "User" WHERE id = $1
`

func (q *Queries) SelectUserByID(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.selectUserByIDStmt, selectUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectUsers = `-- name: SelectUsers :many
SELECT id, name, email, phone, role, created_at, updated_at FROM "User"
`

func (q *Queries) SelectUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.selectUsersStmt, selectUsers)
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
			&i.Role,
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
UPDATE "User" SET name = $1, email = $2, phone = $3, role = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5
`

type UpdateUserParams struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Phone sql.NullString `json:"phone"`
	Role  sql.NullString `json:"role"`
	ID    int32          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Role,
		arg.ID,
	)
	return err
}
