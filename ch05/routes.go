package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HomePage",
		"GET",
		"/",
		HomePage,
	},
	Route{
		"CityList",
		"GET",
		"/city",
		CityList,
	},
	Route{
		"CityDisplay",
		"GET",
		"/city/{id}",
		CityDisplay,
	},
	Route{
		"CityAdd",
		"POST",
		"/cityadd",
		CityAdd,
	},
	Route{
		"CityDelete",
		"GET",
		"/citydel/{id}",
		CityDelete,
	},
}
