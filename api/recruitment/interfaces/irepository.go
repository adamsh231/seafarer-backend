package interfaces

import (
	"seafarer-backend/domain/models"

	"gorm.io/gorm"
)

type IRecruitmentsRepository interface {
	Add(recruitment *models.Recruitments, tx *gorm.DB) (err error)

	UpdateByIDUser(idUser string, model models.Recruitments, tx *gorm.DB) (err error)

	FilterByStatusRecruitment(offset, limit int, orderBy, sort, search, status string) (model []models.RecruitmentsDetail, count int64, err error)
}
