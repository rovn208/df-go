package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBUrl string `mapstructure:"DATABASE_URL"`
	Port  string `mapstructure:"PORT"`
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

func LoadEnv() (config Config, err error) {
	return Config{
		Port:  GetEnvStr("PORT", "8080"),
		DBUrl: GetEnvStr("DATABASE_URL", ""),
	}, nil
}

func GetEnvStr(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
