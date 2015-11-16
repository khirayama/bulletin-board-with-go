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
		message.Index,
	},
	Route{
		"MessageCreate",
		"POST",
		"/messages",
		message.Create,
	},
	Route{
		"MessageEdit",
		"GET",
		"/messages/{id}",
		message.Edit,
	},
	Route{
		"MessageUpdate",
		"PUT",
		"/messages/{id}",
		message.Update,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/messages/{id}",
		message.Delete,
	},
}
