package router

import (
	"github.com/Harduim/the-big-cake/controllers"
	"github.com/gin-gonic/gin"
)

func RouteSport(r *gin.Engine) {
	routeGroup := r.Group("/sports")
	s := controllers.Sport{}
	routeGroup.GET("/", s.GetAll)
	routeGroup.PUT("/:id", s.Update)
}
