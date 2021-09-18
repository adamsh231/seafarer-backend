package models

import "time"

type Admin struct {
	ID        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CompanyID string    `gorm:"column:company_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func NewAdmin() *Admin {
	return &Admin{}
}

func (Admin) TableName() string {
	return "admins"
}
