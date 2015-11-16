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
		messageHandlers.Index,
	},
	Route{
		"MessageCreate",
		"POST",
		"/messages",
		messageHandlers.Create,
	},
	Route{
		"MessageEdit",
		"GET",
		"/messages/{id}",
		messageHandlers.Edit,
	},
	Route{
		"MessageUpdate",
		"PUT",
		"/messages/{id}",
		messageHandlers.Update,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/messages/{id}",
		messageHandlers.Delete,
	},
}
