package cfrouting

import "github.com/gorilla/mux"

// CreateRouter creates a mux.Router for handling http requests
func CreateRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}

	return router
}
