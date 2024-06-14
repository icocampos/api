package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes of the API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequered bool
}

// Configure put all routes into the route
func Config(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
