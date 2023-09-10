package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Harduim/the-big-cake/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	*viper.Viper
	Db *gorm.DB
}

func NewDBConfig() *DBConfig {
	config := &DBConfig{
		Viper: viper.New(),
	}

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}
	host := config.GetString("POSTGRES_HOST")
	port := config.GetInt("POSTGRES_PORT")
	username := config.GetString("POSTGRES_USER")
	password := config.GetString("POSTGRES_PASSWORD")
	database := config.GetString("POSTGRES_DB")

	var db *gorm.DB
	var err error
	dsn := "user=" + username + " password=" + password + " dbname=" + database + " host=" + host + " port=" + strconv.Itoa(port) + " TimeZone=UTC sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&models.User{}, &models.Country{}, &models.Sport{}, &models.Bet{}, &models.SportCategory{})

	if err != nil {
		fmt.Println("Database connection failed:", err.Error())
		os.Exit(2)
	}
	config.Db = db

	return config
}
