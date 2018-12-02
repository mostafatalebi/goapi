package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	routesCol "api/routes"

	"github.com/gorilla/mux"
)

type App struct{}

//RouteItem A single route entity.
type RouteItem struct {
	Methods []string `json:"methods"`
	Handler string   `json:"handler"`
	Name    string
}

func (app *App) loadRoutes() []RouteItem {
	var appRoutes map[string]*RouteItem
	fileBytes, err := ioutil.ReadFile("configs/routes.json")
	if err != nil {
		app.exit("Failed to load routes.json file.")
	}
	err = json.Unmarshal(fileBytes, &appRoutes)
	if err != nil {
		app.exit("Failed to parse contents of routes.json as a JSON object.")
	}

	var rItems = len(appRoutes)
	if rItems < 1 {
		app.exit("routes.json file must contain at least a route")
	}

	routesList := make([]RouteItem, rItems)

	for k, v := range appRoutes {
		_ = v
		fmt.Println(k)
		if len(v.Methods) > 0 {
			v.Name = k
			routesList = append(routesList, *v)
		} else {
			app.exit("No methods specified for " + k)
		}
	}

	return routesList
}

func (app *App) bindRoutes(routes []RouteItem) error {
	routerEngine := mux.NewRouter()
	for _, v := range routes {
		for _, mth := range v.Methods {
			if val, ok := routesCol.Collection[v.Handler]; ok {

			}
			//routerEngine.HandleFunc(v.Name, func(v.Handler)).Methods(mth)
		}
	}
	return nil
}

//Start ...
func Start() {
	app := App{}
	routeColl := app.loadRoutes()
	app.bindRoutes(routeColl)
}

func (app *App) exit(msg string) {
	log.Panic(msg)
	os.Exit(1)
}
