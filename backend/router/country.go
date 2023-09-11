package router

import (
	"github.com/Harduim/the-big-cake/controllers"
	"github.com/gin-gonic/gin"
)

func RouteCountry(r *gin.Engine) {
	routeGroup := r.Group("/countries")
	c := controllers.Country{}
	routeGroup.GET("/", c.GetAll)
	// routeGroup.POST("/", c.Create)
	// routeGroup.PUT("/:id", c.Update)
	// routeGroup.DELETE("/:id", c.Delete)
}
