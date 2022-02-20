package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBname   string
}

type LogConfig struct {
	Level string
}

type RPCConfig struct {
	Endpoint string
}

var (
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
	RPC      RPCConfig
)

func Load() (err error) {
	// read in environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ETH")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".././config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("logger.level", "debug")

	// fmt.Printf("Using configuration file '%s'\n", viper.ConfigFileUsed())

	if err := viper.UnmarshalKey("server", &Server); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("logger", &Log); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("database", &Database); err != nil {
		return err
	}

	if envHost := viper.Get("database.host"); envHost != nil {
		Database.Host = envHost.(string)
	}

	if err := viper.UnmarshalKey("rpc", &RPC); err != nil {
		return err
	}
	return
}
