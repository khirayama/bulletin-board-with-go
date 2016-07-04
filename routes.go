package main

var Routing = Routes{
	Route{
		"MessagesIndex",
		"GET",
		"/messages",
		MessagesIndex,
	},
	Route{
		"MessageCreate",
		"POST",
		"/messages",
		MessageCreate,
	},
	Route{
		"MessageEdit",
		"GET",
		"/messages/{id}",
		MessageEdit,
	},
	Route{
		"MessageUpdate",
		"PUT",
		"/messages/{id}",
		MessageUpdate,
	},
	Route{
		"MessageDelete",
		"DELETE",
		"/messages/{id}",
		MessageDelete,
	},
}
