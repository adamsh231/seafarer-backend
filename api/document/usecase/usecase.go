package usecase

import (
	"html/template"
	"os"
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
	pdfFile, err := libPDF.GeneratePdf(uc.DocAFE.Name, uc.DocAFE.Path, uc.UserID, fMap, afe)
	defer pdfFile.Close()
	if err != nil {
		api.NewErrorLog("AfeUseCase.DownloadAFE", "libPDF.GeneratePdf", err.Error())
		return "", err
	}

	//library function map
	minioLibrary := libraries.MinioLibrary{
		MinioClient: uc.Minio,
		BucketName:  uc.MinioBucketName,
	}

	presignedURL, err := minioLibrary.UploadTypeOsFileWithPresignedKey(pdfFile)

	if err != nil {
		return "", err
	}

	//delete local file
	os.Remove(pdfFile.Name())

	return presignedURL, err
}
