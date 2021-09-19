package interfaces

import "seafarer-backend/domain/models"

type IAFERepository interface {
	IsIDExist(id string) (isExist bool, err error)

	Read(IDUser string) (afe models.AFE, err error)

	Add(afe *models.AFE) (err error)

	Update(afe *models.AFE) (err error)
}
