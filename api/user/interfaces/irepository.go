package interfaces

import (
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Add(user *models.User, tx *gorm.DB) (err error)

	ReadByEmail(email string, user *models.User) (err error)

	UpdateVerifiedByEmail(email string, tx *gorm.DB) (err error)

	UpdatePasswordByEmail(email string, password string, tx *gorm.DB) (err error)

	Filter(offset, limit int, orderBy, sort, search string) (model []models.User, count int64, err error)

	FilterUserAvailable(offset, limit int, orderBy, sort, search string) (model []models.User, count int64, err error)

	FilterByStatusRecruitment(offset, limit int, orderBy, sort, search, status string) (model []models.User, count int64, err error)

	Update(id string, model models.User, tx *gorm.DB) (err error)
}
