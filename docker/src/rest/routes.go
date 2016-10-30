package main

import (
	"github.com/gorilla/mux"
	"net/http"
  "rest/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
		"Register",
		"POST",
		"/api/register",
		handlers.Register,
	},
	Route{
		"Authenticate",
		"POST",
		"/api/authenticate",
    handlers.Authenticate,
	},
  Route{
    "GetUsers",
    "GET",
    "/api/users",
    handlers.Validate(handlers.GetUsers),
  },
}
