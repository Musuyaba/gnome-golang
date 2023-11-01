package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	DATABASE_HOST     string `mapstructure:"DATABASE_HOST"`
	DATABASE_USER     string `mapstructure:"DATABASE_USER"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	DATABASE_PORT     string `mapstructure:"DATABASE_PORT"`
	APPNAME           string `mapstructure:"APPNAME"`
	APIKEY            string `mapstructure:"APIKEY"`
	HOSTNAME          string `mapstructure:"HOSTNAME"`
	PORT              string `mapstructure:"PORT"`
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
