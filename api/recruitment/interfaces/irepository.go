package interfaces

import (
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type IRecruitmentsRepository interface {
	Add(recruitment *models.Recruitments, tx *gorm.DB) (err error)
}
