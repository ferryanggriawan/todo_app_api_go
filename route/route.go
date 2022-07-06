package route

import (
	"todo_app_api_go/app/middleware"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {
	Route = gin.Default()
}

func Setup() *gin.Engine {
	Route.Use(middleware.Logger())

	RootRoute()

	return Route
}
