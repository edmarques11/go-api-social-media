package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent application routes
type Route struct {
	Uri         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Configure applay all routes in *mux.Routes
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, login)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.Uri,
				middlewares.Logger(
					middlewares.Authenticate(route.Function),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
