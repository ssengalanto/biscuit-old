//nolint:gosec //intentional
package constants

import "time"

// Constants for application configurations.
const (
	MaxHeaderBytes      = 1 << 20
	RateLimit           = 100
	IdleTimeout         = time.Minute
	ReadTimeout         = 10 * time.Second
	WriteTimeout        = 30 * time.Second
	Timeout             = time.Minute
	ShutdownGracePeriod = 15 * time.Second
	Dev                 = "development"
	Test                = "test"
	Prod                = "production"
)

// Constants for internal packages.
const (
	ZapLogType      = "zap"
	ViperConfigType = "viper"
	PgsqlDriver     = "pgx"
	SlqMockDriver   = "sqlmock"
)

// Constants for environment variable keys.
const (
	AppName                   = "APP_NAME"
	AppEnv                    = "APP_ENV"
	AccountServicePort        = "ACCOUNT_SERVICE_PORT"
	AuthServicePort           = "AUTH_SERVICE_PORT"
	PgsqlUser                 = "PGSQL_USER"
	PgsqlPassword             = "PGSQL_PASSWORD"
	PgsqlDB                   = "PGSQL_DB"
	PgsqlPort                 = "PGSQL_PORT"
	PgsqlDSN                  = "PGSQL_DSN"
	MongoUser                 = "MONGO_USER"
	MongoPassword             = "MONGO_PASSWORD"
	MongoPort                 = "MONGO_PORT"
	MongoDSN                  = "MONGO_DSN"
	RedisPort                 = "REDIS_PORT"
	RedisURL                  = "REDIS_URL"
	RedisDB                   = "REDIS_DB"
	RedisPassword             = "REDIS_PASSWORD"
	RedisInsightPort          = "REDIS_INSIGHT_PORT"
	RedisInsightURL           = "REDIS_INSIGHT_URL"
	LogType                   = "LOG_TYPE"
	ConfigType                = "CONFIG_TYPE"
	JwtAccessTokenPrivateKey  = "JWT_ACCESS_TOKEN_PRIVATE_KEY"
	JwtAccessTokenPublicKey   = "JWT_ACCESS_TOKEN_PUBLIC_KEY"
	JwtAccessTokenExpiry      = "JWT_ACCESS_TOKEN_EXPIRY"
	JwtRefreshTokenPrivateKey = "JWT_REFRESH_TOKEN_PRIVATE_KEY"
	JwtRefreshTokenPublicKey  = "JWT_REFRESH_TOKEN_PUBLIC_KEY"
	JwtRefreshTokenExpiry     = "JWT_REFRESH_TOKEN_EXPIRY"
)
