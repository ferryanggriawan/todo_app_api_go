package main

import (
	"todo_app_api_go/route"
)

func main() {
	const port string = ":9000"

	app := route.Setup()

	app.Run(port)
}
