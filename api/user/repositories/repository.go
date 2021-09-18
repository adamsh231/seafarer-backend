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

func (repo UserRepository) ReadByEmail(email string, model *models.User) (err error) {
	return repo.Postgres.Where("email = ?", email).First(model).Error
}
