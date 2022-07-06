package route

import (
	"todo_app_api_go/handler"
)

func AuthRoute() {
	auth := Root.Group("/auth")
	{
		auth.POST("/register", handler.RegisterHandler)
	}
}
