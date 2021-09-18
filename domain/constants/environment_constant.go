package constants

const (
	EnvironmentDirectory = "../../.env"

	EnvironmentAppName            = "APP_NAME"
	EnvironmentAppRestPort        = "APP_REST_PORT"
	EnvironmentAppAssetsDirectory = "APP_ASSETS_DIR"

	EnvironmentJWTSecretKey = "JWT_SECRET_KEY"

	EnvironmentAPIKey = "API_KEY"

	EnvironmentLogMode         = "LOG_MODE"
	EnvironmentLogPostgresMode = "LOG_POSTGRES_MODE"

	EnvironmentPostgresMigrationDirectory = "POSTGRES_MIGRATION_DIRECTORY"
	EnvironmentPostgresMigrationDialect   = "POSTGRES_MIGRATION_DIALECT"
	EnvironmentPostgresDBHost             = "POSTGRES_DB_HOST"
	EnvironmentPostgresDBUser             = "POSTGRES_DB_USER"
	EnvironmentPostgresDBPassword         = "POSTGRES_DB_PASSWORD"
	EnvironmentPostgresDBName             = "POSTGRES_DB_NAME"
	EnvironmentPostgresDBPort             = "POSTGRES_DB_PORT"
	EnvironmentPostgresDBSSLMode          = "POSTGRES_DB_SSL_MODE"

	EnvironmentRedisHost     = "REDIS_HOST"
	EnvironmentRedisUser     = "REDIS_USER"
	EnvironmentRedisPassword = "REDIS_PASSWORD"
	EnvironmentRedisDB       = "REDIS_DB"

	EnvironmentSMTPHost       = "SMTP_HOST"
	EnvironmentSMTPPort       = "SMTP_PORT"
	EnvironmentSMTPUser       = "SMTP_USER"
	EnvironmentSMTPPassword   = "SMTP_PASSWORD"

	EnvironmentMongoHost     = "MONGO_HOST"
	EnvironmentMongoPort     = "MONGO_PORT"
	EnvironmentMongoUser     = "MONGO_USER"
	EnvironmentMongoPassword = "MONGO_PASSWORD"
	EnvironmentMongoDatabase = "MONGO_DATABASE"

	EnvironmentMinioHost            = "MINIO_HOST"
	EnvironmentMinioAccessKeyID     = "MINIO_ACCESS_KEY_ID"
	EnvironmentMinioSecretAccessKey = "MINIO_SECRET_ACCESS_KEY"
	EnvironmentMinioBucket          = "MINIO_BUCKET"
	EnvironmentMinioSSL             = "MINIO_SSL"

	EnvironmentEndPointStorage = "ENDPOINT_STORAGE_UPLOAD"

	EnvironmentDocAFEName = "DOC_AFE_NAME"
)
