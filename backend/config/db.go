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
	Host     string `mapstructure:"POSTGRES_HOST" env:"POSTGRES_HOST" default:"localhost"`
	Port     int    `mapstructure:"POSTGRES_PORT" env:"POSTGRES_PORT" default:"5432"`
	Username string `mapstructure:"POSTGRES_USER" env:"POSTGRES_USER" default:"unset"`
	Password string `mapstructure:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" default:"unset"`
	Database string `mapstructure:"POSTGRES_DB" env:"POSTGRES_DB" default:"unset"`
	Db       *gorm.DB
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

	var db *gorm.DB
	var err error
	dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&models.User{}, &models.Country{}, &models.Sport{}, &models.Bet{}, &models.SportCategory{})

	if err != nil {
		fmt.Println("Database connection failed:", err.Error())
		os.Exit(2)
	}
	config.Db = db

	return config
}
