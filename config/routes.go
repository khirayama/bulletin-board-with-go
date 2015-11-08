package config

import (
	"../handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Routing = Routes{
	Route{
		"MessageIndex",
		"GET",
		"/messages",
		handlers.MessageIndex,
	},
	Route{
		"MessageCreate",
		"POST",
		"/messages",
		handlers.MessageCreate,
	},
	Route{
		"MessageEdit",
		"GET",
		"/messages/{id}",
		handlers.MessageEdit,
	},
	Route{
		"MessageUpdate",
		"PUT",
		"/messages/{id}",
		handlers.MessageUpdate,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/messages/{id}",
		handlers.MessageDelete,
	},
}
