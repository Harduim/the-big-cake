package router

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
	RouteUser(r)
	RouteSport(r)
	RouteCountry(r)
	RouteBet(r)
	RouteSportCategory(r)
}
