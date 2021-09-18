package interfaces

import (
	"gorm.io/gorm"
	"seafarer-backend/domain/models"
)

type IUserRepository interface {
	Add(user *models.User, tx *gorm.DB) (err error)

	ReadByEmail(email string, user *models.User) (err error)

	UpdateVerifiedByEmail(email string, tx *gorm.DB) (err error)

	UpdatePasswordByEmail(email string, password string, tx *gorm.DB) (err error)
}
