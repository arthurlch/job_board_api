package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	Phone     string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type JobSeeker struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"unique"`
	Resume    string    `gorm:"type:text"`
	Skills    []string  `gorm:"type:text[]"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Company struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"unique"`
	Name        string    `gorm:"size:255"`
	Email       string    `gorm:"size:255"`
	Phone       string    `gorm:"size:255"`
	Website     string    `gorm:"size:255"`
	Logo        string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Education struct {
	ID            uint      `gorm:"primaryKey"`
	JobSeekerID   uint
	Institution   string    `gorm:"size:255"`
	Degree        string    `gorm:"size:255"`
	FieldOfStudy  string    `gorm:"size:255"`
	StartDate     time.Time `gorm:"type:date"`
	EndDate       time.Time `gorm:"type:date"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Experience struct {
	ID            uint      `gorm:"primaryKey"`
	JobSeekerID   uint
	Title         string    `gorm:"size:255"`
	Company       string    `gorm:"size:255"`
	Location      string    `gorm:"size:255"`
	StartDate     time.Time `gorm:"type:date"`
	EndDate       time.Time `gorm:"type:date"`
	Description   string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Job struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	Requirements string   `gorm:"type:text"`
	Location    string    `gorm:"size:255"`
	Salary      float64   `gorm:"type:decimal"`
	CompanyID   uint
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Application struct {
	ID          uint      `gorm:"primaryKey"`
	JobSeekerID uint
	JobID       uint
	CoverLetter string    `gorm:"type:text"`
	Resume      string    `gorm:"type:text"`
	Status      string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
