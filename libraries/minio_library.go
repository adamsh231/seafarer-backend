package libraries

import (
	"context"
	"mime/multipart"
	"net/url"
	"os"
	"seafarer-backend/domain/constants"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioLibrary struct {
	MinioClient *minio.Client
	BucketName  string
}

func (lib MinioLibrary) Upload(file *multipart.FileHeader) (err error) {

	openFile, err := file.Open()
	if err != nil {
		return err
	}
	defer openFile.Close()

	if _, err = lib.MinioClient.PutObject(context.Background(), lib.BucketName, file.Filename, openFile, file.Size, minio.PutObjectOptions{}); err != nil {
		return err
	}

	return err
}

func (lib MinioLibrary) GetPresignedKey(filename string, expire time.Duration) (urlKey string, err error) {

	// get presigned key
	reqParams := make(url.Values)
	presignedURL, err := lib.MinioClient.PresignedGetObject(context.Background(), lib.BucketName, filename, expire, reqParams)
	if err != nil {
		return urlKey, err
	}

	return presignedURL.String(), err
}

func (lib MinioLibrary) UploadTypeOsFileWithPresignedKey(file *os.File) (presignedURL string, err error) {
	fileStat, err := file.Stat()
	if err != nil {
		return presignedURL, err
	}

	if _, err := lib.MinioClient.PutObject(context.Background(), lib.BucketName, file.Name(), file, fileStat.Size(), minio.PutObjectOptions{}); err != nil {
		return presignedURL, err
	}

	presignedURL, err = lib.GetPresignedKey(file.Name(), constants.MinioDefaultExpiredPresignedKey)
	if err != nil {
		return presignedURL, err
	}

	return presignedURL, err
}
