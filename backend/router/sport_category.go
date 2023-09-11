package router

import (
	"github.com/Harduim/the-big-cake/controllers"
	"github.com/gin-gonic/gin"
)

func RouteSportCategory(r *gin.Engine) {
	routeGroup := r.Group("/sport_categories")
	s := controllers.SportCategory{}
	routeGroup.GET("/", s.GetAll)
}
