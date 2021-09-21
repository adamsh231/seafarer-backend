package domain

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"seafarer-backend/domain/constants"
	"seafarer-backend/libraries"
	"strconv"

	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/minio/minio-go/v7"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/gorm"
)

type Config struct {
	DocAFE             AFEDetail
	App                *fiber.App
	Validator          *validator.Validate
	Postgres           *gorm.DB
	PostgresConnection *sql.DB
	Redis              *redis.Client
	Mail               libraries.MailLibrary
	MongoClient        *mongo.Client
	MongoDatabase      *mongo.Database
	Minio              *minio.Client
	MinioBucketName    string
}

type AFEDetail struct {
	Path string
	Name string
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

	// mongo client
	mongoUser := os.Getenv(constants.EnvironmentMongoUser)
	mongoPass := os.Getenv(constants.EnvironmentMongoPassword)
	mongoHost := os.Getenv(constants.EnvironmentMongoHost)
	mongoPort := os.Getenv(constants.EnvironmentMongoPort)
	mongoDatabase := os.Getenv(constants.EnvironmentMongoDatabase)
	atlasURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPass, mongoHost, mongoPort)
	config.MongoClient, err = mongo.NewClient(options.Client().ApplyURI(atlasURI))
	if err != nil {
		return config, err
	}
	if err = config.MongoClient.Connect(context.Background()); err != nil {
		return config, err
	}
	if err = config.MongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		return config, err
	}
	config.MongoDatabase = config.MongoClient.Database(mongoDatabase)

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

	// minio
	config.MinioBucketName = os.Getenv(constants.EnvironmentMinioBucket)
	minioEndpoint := os.Getenv(constants.EnvironmentMinioHost)
	minioAccessKeyID := os.Getenv(constants.EnvironmentMinioAccessKeyID)
	minioSecretAccessKey := os.Getenv(constants.EnvironmentMinioSecretAccessKey)
	minioSSL := false
	if os.Getenv(constants.EnvironmentMinioSSL) == "1" {
		minioSSL = true
	}
	config.Minio, err = minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, ""),
		Secure: minioSSL,
	})
	if err != nil {
		return config, err
	}
	
	//set doc afe directory
	config.DocAFE.Name = os.Getenv(constants.EnvironmentDocAFEName)
	config.DocAFE.Path = os.Getenv(constants.EnvironmentAppAssetsDirectory) + config.DocAFE.Name

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
