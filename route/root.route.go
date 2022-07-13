package route

import (
	root_handler "todo_app_api_go/handler/root"

	"github.com/gin-gonic/gin"
)

type Root struct {
	Router *gin.RouterGroup
}

func (root *Root) SetupRoute(r *Route) {
	root.Router = r.Router.Group("/api/v1")
	{
		root.Router.GET("/", root_handler.RootHandler)
		UserRoute(root)
		AuthRoute(root)
	}
}
