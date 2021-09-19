package presenters

import (
	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/models"
	"seafarer-backend/libraries"

	"github.com/minio/minio-go/v7"
)

type FileBrowsePresenter struct {
	ListData []interface{} `json:"list"`
}
type fileBrowse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UrlKey string `json:"url_key"`
}

func NewFileBrowsePresenter() FileBrowsePresenter {
	return FileBrowsePresenter{}
}

func (presenter FileBrowsePresenter) Build(files []models.File, minioClient *minio.Client, minioBucket string, total int64) FileBrowsePresenter {
	var fileBrowseData FileBrowsePresenter
	if total > 0 {
		// get presigned key
		minioLibrary := libraries.MinioLibrary{
			MinioClient: minioClient,
			BucketName:  minioBucket,
		}
		for _, val := range files {
			urlKey := ""

			// file name minio
			fileName := val.UserID + "_" + val.Name

			// generate url presignedKey
			key, err := minioLibrary.GetPresignedKey(fileName, constants.MinioDefaultExpiredPresignedKey)
			if err == nil {
				urlKey = key
			}

			fileTemp := fileBrowse{
				ID:     val.ID,
				Name:   val.Name,
				UrlKey: urlKey,
			}
			fileBrowseData.ListData = append(fileBrowseData.ListData, fileTemp)
		}
	} else {
		fileBrowseData.ListData = []interface{}{}
	}
	return fileBrowseData
}
