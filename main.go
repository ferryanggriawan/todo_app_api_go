package main

import (
	_ "todo_app_api_go/app"
	"todo_app_api_go/app/database"
	"todo_app_api_go/route"
)

func main() {
	const port string = ":9000"

	db := &database.Connection{}
	db.Initialize()

	appRoute := &route.Route{}
	appRoute.Initialize()
	appRoute.Setup()

	appRoute.Router.Run(port)
}
