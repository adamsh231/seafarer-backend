package interfaces

import (
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type IFileRepository interface {
	IsExist(id string, name string) (isExist bool, err error)

	Read(id string, file *models.File) (err error)

	Browse(userID string, offset, limit int, search, orderBy, sort string) (files []models.File, count int64, err error)

	Add(user *models.File, tx *gorm.DB) (err error)

	Delete(id string, tx *gorm.DB) (err error)
}
