package route

import (
	"todo_app_api_go/handler"
)

func UserRoute() {
	user := Root.Group("/user")
	{
		user.GET("/", handler.UserHandler)
	}
}
