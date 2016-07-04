package main

import (
	"./handlers"
)

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
