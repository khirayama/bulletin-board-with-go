package config

import (
	"../define"
	"../handlers"
)

var Routing = define.Routes{
	define.Route{
		"MessageIndex",
		"GET",
		"/messages",
		messageHandlers.Index,
	},
	define.Route{
		"MessageCreate",
		"POST",
		"/messages",
		messageHandlers.Create,
	},
	define.Route{
		"MessageEdit",
		"GET",
		"/messages/{id}",
		messageHandlers.Edit,
	},
	define.Route{
		"MessageUpdate",
		"PUT",
		"/messages/{id}",
		messageHandlers.Update,
	},
	define.Route{
		"MessageDelete",
		"DELETE",
		"/messages/{id}",
		messageHandlers.Delete,
	},
}
