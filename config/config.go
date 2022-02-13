package config

import (
	"github.com/spf13/viper"
)

type Env struct {
	// APP
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_Port"`

	// DB
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

// Load initializes the global ENV.
func Load() Env {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("POSTGRES_HOSTNAME", "localhost")
	viper.SetDefault("POSTGRES_USERNAME", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "123")
	viper.SetDefault("POSTGRES_DATABASE", "postgres")
	viper.SetDefault("POSTGRES_PORT", "5432")

	viper.SetDefault("NAME", "XM")
	viper.SetDefault("PORT", "5000")

	settings := Env{
		AppName:    viper.GetString("NAME"),
		AppPort:    viper.GetString("PORT"),
		DBName:     viper.GetString("POSTGRES_DATABASE"),
		DBHost:     viper.GetString("POSTGRES_HOSTNAME"),
		DBPort:     viper.GetString("POSTGRES_PORT"),
		DBUser:     viper.GetString("POSTGRES_USERNAME"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
	}

	return settings
}
