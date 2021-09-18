package repositories

import (
	"gorm.io/gorm"
	"seafarer-backend/api/admin/interfaces"
	"seafarer-backend/domain/models"
)

type AdminRepository struct {
	Postgres *gorm.DB
}

func NewAdminRepository(postgres *gorm.DB) interfaces.IAdminRepository {
	return &AdminRepository{Postgres: postgres}
}

func (repo AdminRepository) ReadByEmail(email string, model *models.Admin) (err error) {
	return repo.Postgres.Where("email = ?", email).First(model).Error
}
