package presenters

import "seafarer-backend/domain/models"

type FileDetailPresenter struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UrlKey string `json:"url_key"`
}

func NewFileDetailPresenter() FileDetailPresenter {
	return FileDetailPresenter{}
}

func (presenter FileDetailPresenter) Build(file *models.File, urlKey string) FileDetailPresenter {
	return FileDetailPresenter{
		ID:     file.UserID,
		Name:   file.Name,
		UrlKey: urlKey,
	}
}
