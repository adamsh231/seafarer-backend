package usecase

import (
	"errors"
	"mime/multipart"
	"seafarer-backend/api"
	"seafarer-backend/api/storage/interfaces"
	repositories "seafarer-backend/api/storage/repository"
	"seafarer-backend/api/storage/router/presenters"
	"seafarer-backend/api/storage/router/requests"
	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/constants/messages"
	"seafarer-backend/domain/models"
	"seafarer-backend/libraries"
	"time"
)

type FileUseCase struct {
	*api.Contract
}

func NewFileUseCase(contract *api.Contract) interfaces.IFileUseCase {
	return &FileUseCase{contract}
}

func (uc FileUseCase) GetWithPresignedKey(id string) (file presenters.FileDetailPresenter, err error) {

	// repo read file
	model := models.NewFile()
	repo := repositories.NewFileRepository(uc.Postgres)
	if err = repo.Read(id, model); err != nil {
		api.NewErrorLog("File.GetPresignedKey", "repo.Read", err.Error())
		return file, err
	}

	// check if file is belong to this user
	if model.UserID != uc.UserID {
		return file, errors.New(messages.UnauthorizedFile)
	}

	// get presigned key
	minioLibrary := libraries.MinioLibrary{
		MinioClient: uc.Minio,
		BucketName:  uc.MinioBucketName,
	}
	// file name minio
	fileName := uc.UserID + "_" + model.Name
	key, err := minioLibrary.GetPresignedKey(fileName, constants.MinioDefaultExpiredPresignedKey)
	if err != nil {
		return
	}

	file = file.Build(model, key)
	return file, err
}

func (uc FileUseCase) Add(file *multipart.FileHeader) (err error) {

	// store file spec to db
	now := time.Now()
	model := &models.File{
		UserID:    uc.UserID,
		Name:      file.Filename,
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo := repositories.NewFileRepository(uc.Postgres)

	// check is name exist
	isExist, err := repo.IsExist(model.UserID, model.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.FileIsExistMessage)
	}

	// upload file
	if err = repo.Add(model, uc.PostgresTX); err != nil {
		api.NewErrorLog("File.Add", "repo.Add", err.Error())
		return err
	}

	// file name
	filename := uc.UserID + "_" + file.Filename

	// minio upload
	minioLibrary := libraries.MinioLibrary{
		MinioClient: uc.Minio,
		BucketName:  uc.MinioBucketName,
	}
	file.Filename = filename
	if err := minioLibrary.Upload(file); err != nil {
		api.NewErrorLog("File.Add", "minioLibrary.Upload", err.Error())
		return err
	}

	return err
}

func (uc FileUseCase) BrowseWithPresignedKey(request *requests.BrowseFilesRequest) (file presenters.FileBrowsePresenter, meta api.MetaResponsePresenter, err error) {

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(request.Page, request.PerPage, request.Order, request.Sort)

	// repo browse files
	repo := repositories.NewFileRepository(uc.Postgres)
	dataFiles, count, err := repo.Browse(uc.UserID, offset, limit, request.Search, orderBy, sort)
	if err != nil {
		api.NewErrorLog("File.BrowsePresignedKey", "repo.Browse", err.Error())
		return file, meta, err
	}

	file = file.Build(dataFiles, uc.Minio, uc.MinioBucketName, count)
	meta = uc.SetPaginationResponse(page, limit, int(count))
	return file, meta, err
}

func (uc FileUseCase) Delete(id string) (err error) {

	// cek file is exist
	model := models.NewFile()
	repo := repositories.NewFileRepository(uc.Postgres)
	if err = repo.Read(id, model); err != nil {
		api.NewErrorLog("File.Delete", "repo.Read", err.Error())
		return err
	}

	// check if file is belong to this user
	if model.UserID != uc.UserID {
		return errors.New(messages.UnauthorizedFile)
	}

	//delete file
	if err = repo.Delete(id, uc.PostgresTX); err != nil {
		api.NewErrorLog("File.Delete", "repo.Delete", err.Error())
		return err
	}

	return err
}

func (uc FileUseCase) PrivateUploadAndGetPresignedKey(file *multipart.FileHeader) (urlPresigned string, err error) {
	// file name
	filename := uc.UserID + "_AFE.pdf"

	// minio upload
	minioLibrary := libraries.MinioLibrary{
		MinioClient: uc.Minio,
		BucketName:  uc.MinioBucketName,
	}
	file.Filename = filename
	if err := minioLibrary.Upload(file); err != nil {
		api.NewErrorLog("File.Upload", "minioLibrary.Upload", err.Error())
		return urlPresigned, err
	}
	urlPresigned, err = minioLibrary.GetPresignedKey(filename, constants.MinioDefaultExpiredPresignedKey)
	if err != nil {
		api.NewErrorLog("File.UploadAndGetPresignedKey", "libraries.MinioLibrary.GetPresignedKey", err.Error())
		return "", err
	}
	return urlPresigned, err
}

func (uc FileUseCase) BrowseByUserID(userID string, request *requests.BrowseFilesRequest) (file presenters.FileBrowsePresenter, meta api.MetaResponsePresenter, err error) {

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(request.Page, request.PerPage, request.Order, request.Sort)

	// repo browse files
	repo := repositories.NewFileRepository(uc.Postgres)
	dataFiles, count, err := repo.Browse(userID, offset, limit, request.Search, orderBy, sort)
	if err != nil {
		api.NewErrorLog("File.BrowseByUserID", "repo.Browse", err.Error())
		return file, meta, err
	}

	file = file.Build(dataFiles, uc.Minio, uc.MinioBucketName, count)
	meta = uc.SetPaginationResponse(page, limit, int(count))
	return file, meta, err
}
