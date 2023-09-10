package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DBConfig struct {
	*viper.Viper
}

func New() *DBConfig {
	config := &DBConfig{
		Viper: viper.New(),
	}

	config.setDefaults()
	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	return config
}

func (config *DBConfig) setDefaults() {
	// Set default App configuration
	config.SetDefault("API_PORT", "3000")

	// Set default database configuration
	config.SetDefault("POSTGRES_DB", "cupcake")
	config.SetDefault("POSTGRES_USER", "USER_UNSET")
	config.SetDefault("POSTGRES_PASSWORD", "PASSWORD_UNSET")
	config.SetDefault("POSTGRES_PORT", 5432)
}
