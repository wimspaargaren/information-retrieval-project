package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route route struct containing the name of the rout, the http method, the pattern and the function which handles the route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes the list of routes
type Routes []Route

//NewRouter initializes the mux with the defined routes.
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetPoints",
		"GET",
		"/getPoints",
		GetPoints,
	},
}
