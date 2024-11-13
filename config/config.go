package config

import (
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	DbDriver             string        `mapstructure:"DB_DRIVER"`
	DbSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"ADDRESS_SERVER"`
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATIOn"`
}

var conf Configuration

func LoadConfig(path string) (config Configuration, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return conf, err
	}
	err = viper.Unmarshal(&conf)
	return conf, err
}

func GetCofig() Configuration {
	return conf
}
