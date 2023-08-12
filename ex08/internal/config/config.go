package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_URL string `mapstructure:"DB_URL"`
	Port   string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
