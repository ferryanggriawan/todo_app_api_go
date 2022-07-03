package route

import (
	"todo_app_api_go/app/handler"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", handler.RootHandler)

	return r
}
