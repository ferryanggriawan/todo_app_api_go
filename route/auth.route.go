package route

import (
	auth_handler "todo_app_api_go/handler/auth"
)

func AuthRoute() {
	auth := Root.Group("/auth")
	{
		auth.POST("/register", auth_handler.RegisterUser)
	}
}
