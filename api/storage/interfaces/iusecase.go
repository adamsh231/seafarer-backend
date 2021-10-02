package interfaces

import (
	"mime/multipart"
	"seafarer-backend/api"
	"seafarer-backend/api/storage/router/presenters"
	"seafarer-backend/api/storage/router/requests"
)

type IFileUseCase interface {
	GetWithPresignedKey(id string) (file presenters.FileDetailPresenter, err error)

	BrowseWithPresignedKey(request *requests.BrowseFilesRequest) (file presenters.FileBrowsePresenter, meta api.MetaResponsePresenter, err error)

	Add(file *multipart.FileHeader) (err error)

	Delete(fileID string) (err error)

	PrivateUploadAndGetPresignedKey(file *multipart.FileHeader) (urlPresigned string, err error)

	BrowseByUserID(userID string, request *requests.BrowseFilesRequest) (file presenters.FileBrowsePresenter, meta api.MetaResponsePresenter, err error)
}
