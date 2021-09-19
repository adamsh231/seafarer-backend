package usecase

import (
	"html/template"
	"seafarer-backend/api"
	"seafarer-backend/api/document/interfaces"
	"seafarer-backend/api/document/repositories"
	"seafarer-backend/domain/models"
	"seafarer-backend/libraries"
)

type AfeUseCase struct {
	*api.Contract
}

func NewAFEUseCase(contract *api.Contract) interfaces.IAFEUseCase {
	return &AfeUseCase{contract}
}

func (uc AfeUseCase) Save(afe *models.AFE) (err error) {

	// init user
	afe.UserID = uc.UserID
	repo := repositories.NewAFERepository(uc.MongoDatabase)

	// check is already fill
	isExist, err := repo.IsIDExist(afe.UserID)
	if err != nil {
		api.NewErrorLog("AfeUseCase.Save", "repo.IsIDExist", err.Error())
		return err
	}
	if isExist {
		if err = repo.Update(afe); err != nil {
			api.NewErrorLog("AfeUseCase.Save", "repo.Update", err.Error())
			return err
		}
	} else {
		if err = repo.Add(afe); err != nil {
			api.NewErrorLog("AfeUseCase.Save", "repo.Add", err.Error())
			return err
		}
	}

	return err
}

func (uc AfeUseCase) Get() (afe models.AFE, err error) {
	//repo read file
	repo := repositories.NewAFERepository(uc.MongoDatabase)
	afe, err = repo.Read(uc.UserID)
	if err != nil {
		api.NewErrorLog("AfeUseCase.Get", "repo.Read", err.Error())
		return afe, err
	}

	return afe, err
}

func (uc AfeUseCase) DownloadAFE() (url string, err error) {
	var afe models.AFE

	//repo read file
	repo := repositories.NewAFERepository(uc.MongoDatabase)
	afe, err = repo.Read(uc.UserID)
	if err != nil {
		api.NewErrorLog("AfeUseCase.DownloadAFE", "repo.Read", err.Error())
		return "", err
	}

	//library function map
	libFuncMap := libraries.NewFuncMapLibrary()

	//library function map
	libPDF := libraries.NewHtmlToPdfLibrary()

	//set fungsi yang akan di gunakan
	fMap := template.FuncMap{
		"increment": libFuncMap.Add,
		"unitconv":  libFuncMap.ConvertUnit,
	}

	//generate file pdf
	_, err = libPDF.GeneratePdf(uc.DocAFE.Name, uc.DocAFE.Path, uc.UserID, fMap, afe)
	if err != nil {
		api.NewErrorLog("AfeUseCase.DownloadAFE", "libraries.NewHtmlToPdfLibrary", err.Error())
		return "", err
	}

	//url, err = uc.uploadToStorageAndGetPresignedKey(uc.UserID)
	//if err != nil {
	//	api.NewErrorLog("AfeUseCase.DownloadAFE", "libraries.MinioLibrary.UploadTypeOsFileWithPresignedKey", err.Error())
	//	return "", err
	//}
	return url, err
}
