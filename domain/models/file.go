package models

import "time"

type File struct {
	ID        string    `gorm:"id"`
	UserID    string    `gorm:"user_id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at"`
}

func NewFile() *File {
	return &File{}
}

func (File) TableName() string {
	return "files"
}
