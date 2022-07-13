package route

import (
	user_handler "todo_app_api_go/handler/user"
)

func UserRoute(r *Root) {
	user := r.Router.Group("/user")
	{
		user.GET("/", user_handler.GetUsers)
	}
}
