package route

import (
	"todo_app_api_go/handler"
)

func Root() {
	api := Route.Group("/api/v1")
	{
		api.GET("/", handler.RootHandler)
	}
}
