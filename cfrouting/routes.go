package cfrouting

import (
	"net/http"
)

// Route represents a route to a function from a http request
type Route struct {
	Name    string           // The name of the route.
	Method  string           // The method to use (GET, PUT, POST, etc)
	Pattern string           // The pattern to get the the route (/foo/bar/{baz})
	Handler http.HandlerFunc // The function to handle the route.
}

// Routes is a collection of routes to use to create a router.
type Routes []Route
