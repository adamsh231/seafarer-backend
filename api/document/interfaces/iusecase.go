package interfaces

import "seafarer-backend/domain/models"

type IAFEUseCase interface {
	Save(afe *models.AFE) (err error)
	Get() (afe models.AFE, err error)
	DownloadAFE() (url string, err error)
}
