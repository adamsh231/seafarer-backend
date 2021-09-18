package repositories

import (
	"gorm.io/gorm"
	"seafarer-backend/api/user/interfaces"
	"seafarer-backend/domain/models"
)

type UserRepository struct {
	Postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) interfaces.IUserRepository {
	return &UserRepository{Postgres: postgres}
}

func (repo UserRepository) Add(model *models.User, tx *gorm.DB) (err error) {
	return tx.Omit("deleted_at").Create(model).Error
}

func (repo UserRepository) ReadByEmail(email string, model *models.User) (err error) {
	return repo.Postgres.Where("email = ?", email).First(model).Error
}

func (repo UserRepository) UpdateVerifiedByEmail(email string, tx *gorm.DB) (err error) {
	return tx.Model(models.NewUser()).Where("email = ?", email).Update("is_verified", true).Error
}

func (repo UserRepository) UpdatePasswordByEmail(email string, password string, tx *gorm.DB) (err error) {
	return tx.Model(models.NewUser()).Where("email = ?", email).Update("password", password).Error
}