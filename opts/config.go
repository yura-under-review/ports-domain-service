package opts

import (
	"github.com/spf13/viper"
	"github.com/yura-under-review/ports-domain-service/grpc"
	"github.com/yura-under-review/ports-domain-service/repository"
)

const (
	defaultEnvPrefix = "APP"

	defaultLogLevel = "DEBUG"

	keyLogLevel   = "LOG_LEVEL"
	keyPrettyLogs = "PRETTY_LOGS"

	keyGprcAddress = "GRPC_ADDRESS"

	keyPostgresUser           = "POSTGRES_USER"
	keyPostgresPassword       = "POSTGRES_PASSWORD"
	keyPostgresHost           = "POSTGRES_HOST"
	keyPostgresPort           = "POSTGRES_PORT"
	keyPostgresDatabaseName   = "POSTGRES_DB_NAME"
	keyPostgresMaxConnections = "POSTGRES_MAX_CONN"
	keyPostgresSSL            = "POSTGRES_SSL"
)

type Config struct {
	LogLevel   string
	PrettyLogs bool

	GRPC     grpc.Config
	Postgres repository.Config
}

func LoadConfigFromEnv() Config {

	viper.AutomaticEnv()

	viper.SetEnvPrefix(defaultEnvPrefix)

	viper.SetDefault(keyLogLevel, defaultLogLevel)
	viper.SetDefault(keyPrettyLogs, false)

	viper.SetDefault(keyGprcAddress, ":8080")

	return Config{
		LogLevel:   viper.GetString(keyLogLevel),
		PrettyLogs: viper.GetBool(keyPrettyLogs),

		GRPC: grpc.Config{
			Address: viper.GetString(keyGprcAddress),
		},

		Postgres: repository.Config{
			User:           viper.GetString(keyPostgresUser),
			Password:       viper.GetString(keyPostgresPassword),
			Host:           viper.GetString(keyPostgresHost),
			Port:           viper.GetInt(keyPostgresPort),
			DatabaseName:   viper.GetString(keyPostgresDatabaseName),
			MaxConnections: viper.GetInt(keyPostgresMaxConnections),
			SSL:            viper.GetBool(keyPostgresSSL),
		},
	}
}
