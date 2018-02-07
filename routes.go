package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defind route structure w/ name, method, pattern, and handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of Route
type Routes []Route

// NewRouter returns *mux.Router w/ customised routes
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
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"Get",
		"/todos/{todoID}",
		TodoShow,
	},
}
