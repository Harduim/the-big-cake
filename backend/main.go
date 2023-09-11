package main

import (
	"github.com/Harduim/the-big-cake/config"
	"github.com/Harduim/the-big-cake/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	router.BindRoutes(r)

	rConfig := config.NewAPIConfig()
	port := rConfig.GetString("API_PORT")

	r.Run(":" + port)

}
