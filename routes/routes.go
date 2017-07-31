package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route : todo
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : todo
type Routes []Route

// NewRouter : todo
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
		"NewsIndex",
		"GET",
		"/news",
		NewsIndex,
	},
	Route{
		"NewsShow",
		"GET",
		"/news/{newsID}",
		NewsShow,
	},
	Route{
		"NewsCreate",
		"Post",
		"/news",
		NewsCreate,
	},
}
