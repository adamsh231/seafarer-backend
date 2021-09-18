package domain

import (
	"context"
	"database/sql"
	"os"
	"seafarer-backend/domain/constants"
	"seafarer-backend/libraries"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Config struct {
	App                *fiber.App
	Validator          *validator.Validate
	Postgres           *gorm.DB
	PostgresConnection *sql.DB
	Redis              *redis.Client
	Mail               libraries.MailLibrary
}

func LoadConfiguration() (config Config, err error) {

	// load env
	if err = godotenv.Load(constants.EnvironmentDirectory); err != nil {
		return config, err
	}

	// logs formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// postgres
	dbLogMode, err := strconv.Atoi(os.Getenv(constants.EnvironmentLogPostgresMode))
	if err != nil {
		return config, err
	}
	postgresLibrary := libraries.PostgresLibrary{
		MigrationDirectory: os.Getenv(constants.EnvironmentPostgresMigrationDirectory),
		MigrationDialect:   os.Getenv(constants.EnvironmentPostgresMigrationDialect),
		DBHost:             os.Getenv(constants.EnvironmentPostgresDBHost),
		DBUser:             os.Getenv(constants.EnvironmentPostgresDBUser),
		DBPassword:         os.Getenv(constants.EnvironmentPostgresDBPassword),
		DBPort:             os.Getenv(constants.EnvironmentPostgresDBPort),
		DBName:             os.Getenv(constants.EnvironmentPostgresDBName),
		DBSSLMode:          os.Getenv(constants.EnvironmentPostgresDBSSLMode),
		LogMode:            dbLogMode,
	}
	config.Postgres, config.PostgresConnection, err = postgresLibrary.ConnectAndValidate()
	if err != nil {
		return config, err
	}
	if err = postgresLibrary.Migrate(config.PostgresConnection); err != nil {
		return config, err
	}

	// redis
	redisDB, err := strconv.Atoi(os.Getenv(constants.EnvironmentRedisDB))
	if err != nil {
		return config, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(constants.EnvironmentRedisHost),
		Username: os.Getenv(constants.EnvironmentRedisUser),
		Password: os.Getenv(constants.EnvironmentRedisPassword),
		DB:       redisDB,
	})
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		return config, err
	}
	config.Redis = rdb

	// mail
	smtpPort, err := strconv.Atoi(os.Getenv(constants.EnvironmentSMTPPort))
	if err != nil {
		return config, err
	}
	config.Mail = libraries.MailLibrary{
		MailHost:     os.Getenv(constants.EnvironmentSMTPHost),
		MailPort:     smtpPort,
		MailUser:     os.Getenv(constants.EnvironmentSMTPUser),
		MailPassword: os.Getenv(constants.EnvironmentSMTPPassword),
	}

	// validator
	config.Validator = validator.New()

	// fiber
	config.App = fiber.New()

	return config, err
}

func HttpListen(app *fiber.App) (err error) {
	if err := app.Listen(os.Getenv(constants.EnvironmentAppRestPort)); err != nil {
		return err
	}

	return err
}
