package route

import (
	auth_handler "todo_app_api_go/handler/auth"
)

func AuthRoute(r *Root) {
	auth := r.Router.Group("/auth")
	{
		auth.POST("/register", auth_handler.RegisterUser)
	}
}
