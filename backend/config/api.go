package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type APIConfig struct {
	*viper.Viper
}

func NewAPIConfig() *APIConfig {
	config := &APIConfig{
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

func (config *APIConfig) setDefaults() {
	// Set default App configuration
	config.SetDefault("API_PORT", 3000)
}
