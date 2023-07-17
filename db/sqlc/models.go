// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
)

type Application struct {
	ID          int32          `json:"id"`
	JobSeekerID sql.NullInt32  `json:"job_seeker_id"`
	JobID       sql.NullInt32  `json:"job_id"`
	CoverLetter sql.NullString `json:"cover_letter"`
	Resume      sql.NullString `json:"resume"`
	Status      sql.NullString `json:"status"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Company struct {
	ID          int32          `json:"id"`
	UserID      sql.NullInt32  `json:"user_id"`
	Name        sql.NullString `json:"name"`
	Email       sql.NullString `json:"email"`
	Phone       sql.NullString `json:"phone"`
	Website     sql.NullString `json:"website"`
	Logo        sql.NullString `json:"logo"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Education struct {
	ID           int32          `json:"id"`
	JobSeekerID  sql.NullInt32  `json:"job_seeker_id"`
	Institution  sql.NullString `json:"institution"`
	Degree       sql.NullString `json:"degree"`
	FieldOfStudy sql.NullString `json:"field_of_study"`
	StartDate    sql.NullTime   `json:"start_date"`
	EndDate      sql.NullTime   `json:"end_date"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type Experience struct {
	ID          int32          `json:"id"`
	JobSeekerID sql.NullInt32  `json:"job_seeker_id"`
	Title       sql.NullString `json:"title"`
	Company     sql.NullString `json:"company"`
	Location    sql.NullString `json:"location"`
	StartDate   sql.NullTime   `json:"start_date"`
	EndDate     sql.NullTime   `json:"end_date"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Job struct {
	ID           int32          `json:"id"`
	Title        sql.NullString `json:"title"`
	Description  sql.NullString `json:"description"`
	Requirements sql.NullString `json:"requirements"`
	Location     sql.NullString `json:"location"`
	Salary       sql.NullInt32  `json:"salary"`
	CompanyID    sql.NullInt32  `json:"company_id"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type Jobseeker struct {
	ID        int32          `json:"id"`
	UserID    sql.NullInt32  `json:"user_id"`
	Resume    sql.NullString `json:"resume"`
	Skills    []string       `json:"skills"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type User struct {
	ID        int32          `json:"id"`
	Name      sql.NullString `json:"name"`
	Email     sql.NullString `json:"email"`
	Phone     sql.NullString `json:"phone"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
