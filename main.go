package main

func main() {
	var app App
	app.Init()
	methods := []string{"GET"}
	ctrl := UserController{}
	app.Router.addRouteItem("/user/add", ctrl.postAdd, methods)
	app.Listen(":6002")
}
