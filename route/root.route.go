package route

import (
	root_handler "todo_app_api_go/handler/root"

	"github.com/gin-gonic/gin"
)

var Root *gin.RouterGroup

func RootRoute() {
	Root = Route.Group("/api/v1")
	{
		Root.GET("/", root_handler.RootHandler)
		UserRoute()
		AuthRoute()
	}
}
