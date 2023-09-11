package router

import (
	"github.com/Harduim/the-big-cake/controllers"
	"github.com/gin-gonic/gin"
)

func RouteBet(r *gin.Engine) {
	routeGroup := r.Group("/bets")
	b := controllers.Bet{}
	routeGroup.GET("/", b.GetAll)
	routeGroup.POST("/", b.Create)
	routeGroup.PUT("/:id", b.Update)
	routeGroup.DELETE("/:id", b.Delete)
}
