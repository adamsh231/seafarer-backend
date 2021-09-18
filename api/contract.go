package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"seafarer-backend/libraries"
)

type Contract struct {
	UserID         string
	UserName       string
	UserEmail      string
	UserIsVerified bool
	UserIsAdmin    bool
	App            *fiber.App
	Validator      *validator.Validate
	Postgres       *gorm.DB
	PostgresTX     *gorm.DB
	Mail           libraries.MailLibrary
	Redis          *redis.Client
}

func NewErrorLog(useCase, specification, message string) {
	logrus.WithFields(logrus.Fields{
		"use_case":      useCase,
		"specification": specification,
	}).Error(message)
}

