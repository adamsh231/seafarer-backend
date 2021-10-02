package repositories

import (
	"seafarer-backend/api/recruitment/interfaces"
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type RecruitmentsRepository struct {
	Postgres *gorm.DB
}

func NewRecruitmentsRepository(postgres *gorm.DB) interfaces.IRecruitmentsRepository {
	return &RecruitmentsRepository{Postgres: postgres}
}

func (repo RecruitmentsRepository) Add(model *models.Recruitments, tx *gorm.DB) (err error) {
	return tx.Omit("deleted_at").Create(model).Error
}
