package main

import (
	_ "todo_app_api_go/app/database"
	"todo_app_api_go/route"
)

func main() {
	const port string = ":9000"

	appRoute := &route.Route{}
	appRoute.Initialize()

	defer recover()
	err := appRoute.Router.Run(port)
	if err != nil {
		panic("App Terminated")
	}
}
