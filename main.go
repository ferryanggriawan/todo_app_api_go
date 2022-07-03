package main

import (
	_ "todo_app_api_go/app/database"
	"todo_app_api_go/route"
)

func main() {
	const port string = ":9000"

	appRoute := route.Setup()

	appRoute.Run(port)
}
