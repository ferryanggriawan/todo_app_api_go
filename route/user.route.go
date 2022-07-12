package route

import (
	user_handler "todo_app_api_go/handler/user"
)

func UserRoute() {
	user := Root.Group("/user")
	{
		user.GET("/", user_handler.GetUsers)
	}
}
