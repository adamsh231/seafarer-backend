package models

import "time"

type User struct {
	ID            string    `gorm:"column:id"`
	Name          string    `gorm:"column:name"`
	Email         string    `gorm:"column:email"`
	Password      string    `gorm:"column:password"`
	IsVerified    bool      `gorm:"column:is_verified"`
	CompanyID     string    `gorm:"column:company_id"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
	RecruitmentID string    `gorm:"column:recruitment_id"`
}

func NewUser() *User {
	return &User{}
}

func (User) TableName() string {
	return "users"
}
