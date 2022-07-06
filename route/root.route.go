package route

import (
	"todo_app_api_go/handler"

	"github.com/gin-gonic/gin"
)

var Root *gin.RouterGroup

func RootRoute() {
	Root = Route.Group("/api/v1")
	{
		Root.GET("/", handler.RootHandler)
		UserRoute()
		AuthRoute()
	}
}
