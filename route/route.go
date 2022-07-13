package route

import (
	"todo_app_api_go/app/middleware"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Router *gin.Engine
}

func (r *Route) Initialize() {
	r.Router = gin.Default()
}

func (r *Route) Setup() *gin.Engine {
	root := &Root{}

	r.Router.Use(middleware.Logger())
	root.SetupRoute(r)

	return r.Router
}
