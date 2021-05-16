package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandler interface {
	HandleRoutes(http.ResponseWriter, *http.Request)
}

func BuildRouter(controllers map[string]RouteHandler) *mux.Router {
	// init router
	r := mux.NewRouter()
	r.StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()

	// api versioning
	v1r := api.PathPrefix("/v1").Subrouter()

	// artists routes
	ar := v1r.PathPrefix("/artists").Subrouter()
	ar.HandleFunc("/", controllers["artists"].HandleRoutes).Methods(http.MethodGet)
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodGet)
	ar.HandleFunc("/", controllers["artists"].HandleRoutes).Methods(http.MethodPost)
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodPut)
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodDelete)

	return r
}
