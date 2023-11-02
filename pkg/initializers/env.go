package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	APPNAME                      string `mapstructure:"APPNAME"`
	APIKEY                       string `mapstructure:"APIKEY"`
	HOSTNAME                     string `mapstructure:"HOSTNAME"`
	PORT                         string `mapstructure:"PORT"`
	DATABASE_USED                string `mapstructure:"DATABASE_USED"`
	DATABASE_POSTGRESQL_NAME     string `mapstructure:"DATABASE_POSTGRESQL_NAME"`
	DATABASE_POSTGRESQL_HOST     string `mapstructure:"DATABASE_POSTGRESQL_HOST"`
	DATABASE_POSTGRESQL_USER     string `mapstructure:"DATABASE_POSTGRESQL_USER"`
	DATABASE_POSTGRESQL_PASSWORD string `mapstructure:"DATABASE_POSTGRESQL_PASSWORD"`
	DATABASE_POSTGRESQL_PORT     string `mapstructure:"DATABASE_POSTGRESQL_PORT"`
	DATABASE_MYSQL_NAME          string `mapstructure:"DATABASE_MYSQL_NAME"`
	DATABASE_MYSQL_HOST          string `mapstructure:"DATABASE_MYSQL_HOST"`
	DATABASE_MYSQL_USER          string `mapstructure:"DATABASE_MYSQL_USER"`
	DATABASE_MYSQL_PASSWORD      string `mapstructure:"DATABASE_MYSQL_PASSWORD"`
	DATABASE_MYSQL_PORT          string `mapstructure:"DATABASE_MYSQL_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
