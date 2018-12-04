package main

import(
	"fmt"
	"os"
)

func main() {
	var app App
	app.Init()
	methods := []string{"POST"}
	ctrl := UserController{}
	app.Router.addRouteItem("/user/add", ctrl.postAdd, methods)
	fmt.Println("Listening...")
	err := app.Listen(":6002")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
