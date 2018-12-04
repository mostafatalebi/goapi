package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TRoutesCollection map[string]RouteItem

type RouteEngine struct {
	RouteColl  TRoutesCollection
	HttpRoutes *mux.Router
}

//RouteItem A single route entity.
type RouteItem struct {
	Methods []string
	Handler func(http.ResponseWriter,
		*http.Request)
	Path string
}

func (rt *RouteEngine) addRouteItem(path string, handler func(http.ResponseWriter,
	*http.Request), methods []string) error {

	var rItem = RouteItem{
		Methods: methods,
		Handler: handler,
		Path:    path,
	}

	if rt.RouteColl == nil {
		rt.RouteColl = make(TRoutesCollection)
	}

	rt.RouteColl[path] = rItem

	return nil
}
