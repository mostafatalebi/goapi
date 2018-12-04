package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//App ...
type App struct {
	Router RouteEngine
}

//deprecated
// func (app *App) loadRoutes() []RouteItem {
// 	var appRoutes map[string]*RouteItem
// 	fileBytes, err := ioutil.ReadFile("configs/routes.json")
// 	if err != nil {
// 		app.exit("Failed to load routes.json file.")
// 	}
// 	err = json.Unmarshal(fileBytes, &appRoutes)
// 	if err != nil {
// 		app.exit("Failed to parse contents of routes.json as a JSON object.")
// 	}

// 	var rItems = len(appRoutes)
// 	if rItems < 1 {
// 		app.exit("routes.json file must contain at least a route")
// 	}

// 	routesList := make([]RouteItem, rItems)

// 	for k, v := range appRoutes {
// 		_ = v
// 		fmt.Println(k)
// 		if len(v.Methods) > 0 {
// 			v.Name = k
// 			routesList = append(routesList, *v)
// 		} else {
// 			app.exit("No methods specified for " + k)
// 		}
// 	}

// 	return routesList
// }

func (app *App) bindRoutes(routes TRoutesCollection) error {
	routerEngine := mux.NewRouter()
	for _, v := range routes {
		for _, mth := range v.Methods {
			routerEngine.HandleFunc(v.Path, v.Handler).Methods(mth)
		}
	}
	app.Router.HttpRoutes = routerEngine
	return nil
}

//AddRoute adds a aroute to the collection of routes
func (app *App) AddRoute(path string, handler func(http.ResponseWriter, *http.Request), methods []string) error {
	return app.Router.addRouteItem(path, handler, methods)
}

//Init Initializes the object
func (app *App) Init() {
	app.Router = RouteEngine{}
}

//Listen starts the server
func (app *App) Listen(addr string) {
	app.bindRoutes(app.Router.RouteColl)
	http.ListenAndServe(addr, app.Router.HttpRoutes)
}

func (app *App) exit(msg string) {
	log.Panic(msg)
	os.Exit(1)
}
