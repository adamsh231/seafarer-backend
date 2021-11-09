package models

import "time"

type Recruitments struct {
	ID           string    `gorm:"column:id"`
	UserID       string    `gorm:"column:user_id"`
	ExpectSalary float64   `gorm:"column:expect_salary"`
	Salary       float64   `gorm:"column:salary"`
	Position     string    `gorm:"column:position"`
	AvailableOn  time.Time `gorm:"column:available_on"`
	SignOn       time.Time `gorm:"column:sign_on"`
	Ship         string    `gorm:"column:ship"`
	Letter       string    `gorm:"column:letter"`
	Status       string    `gorm:"column:status"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at"`
}

func NewRecruitments() *Recruitments {
	return &Recruitments{}
}

func (Recruitments) TableName() string {
	return "recruitments"
}

type RecruitmentsDetail struct {
	ID           string    `gorm:"column:id"`
	UserID       string    `gorm:"column:user_id"`
	UserName     string    `gorm:"column:name"`
	ExpectSalary float64   `gorm:"column:expect_salary"`
	Salary       float64   `gorm:"column:salary"`
	Position     string    `gorm:"column:position"`
	AvailableOn  time.Time `gorm:"column:available_on"`
	SignOn       time.Time `gorm:"column:sign_on"`
	Ship         string    `gorm:"column:ship"`
	Letter       string    `gorm:"column:letter"`
	Status       string    `gorm:"column:status"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at"`
	IsFailed     bool      `gorm:"column:is_failed"`
}
