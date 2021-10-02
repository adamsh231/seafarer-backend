package models

import "time"

type Recruitments struct {
	ID          string    `gorm:"column:id"`
	UserID      string    `gorm:"column:user_id"`
	Salary      string    `gorm:"column:salary"`
	Position    string    `gorm:"column:position"`
	AvailableOn time.Time `gorm:"column:available_on"`
	SignOn      time.Time `gorm:"column:sign_on"`
	Ship        string    `gorm:"column:ship"`
	Letter      string    `gorm:"column:letter"`
	Status      string    `gorm:"column:status"`
	CompanyID   string    `gorm:"column:company_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
}
