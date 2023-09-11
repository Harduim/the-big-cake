package router

import (
	"github.com/Harduim/the-big-cake/controllers"
	"github.com/gin-gonic/gin"
)

func RouteUser(r *gin.Engine) {
	routeGroup := r.Group("/users")
	u := controllers.User{}
	routeGroup.GET("/", u.GetAll)
}
