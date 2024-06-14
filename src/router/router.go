package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate the return one router with configured routes
func Generator() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
