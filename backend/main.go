package main

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/models"
)

func main() {
	dbConfig := config.NewDBConfig()
	db := dbConfig.Db
	db.Find(&models.User{})
}
