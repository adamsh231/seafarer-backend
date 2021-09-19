package api

import (
	"seafarer-backend/libraries"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Contract struct {
	UserID          string
	UserName        string
	UserEmail       string
	UserIsVerified  bool
	UserIsAdmin     bool
	App             *fiber.App
	Validator       *validator.Validate
	Postgres        *gorm.DB
	PostgresTX      *gorm.DB
	Mail            libraries.MailLibrary
	Redis           *redis.Client
	MongoDatabase   *mongo.Database
	DocAFE          AFEDetail
	Minio           *minio.Client
	MinioBucketName string
}

type AFEDetail struct {
	Path string
	Name string
}

func NewErrorLog(useCase, specification, message string) {
	logrus.WithFields(logrus.Fields{
		"use_case":      useCase,
		"specification": specification,
	}).Error(message)
}

const (
	//default limit for pagination
	defaultLimit = 10

	//max limit for pagination
	maxLimit = 50

	//default order by
	defaultOrderBy = "created_at"

	//default sort
	defaultSort = "asc"

	//default last page for pagination
	defaultLastPage = 0
)

func (uc Contract) SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int) (res MetaResponsePresenter) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	res = MetaResponsePresenter{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		LastPage:    lastPage,
	}

	return res
}
