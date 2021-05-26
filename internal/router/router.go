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
	ar.HandleFunc("/", controllers["artists"].HandleRoutes).Methods(http.MethodPost).Name("createArtist")
	ar.HandleFunc("/errtest", controllers["artists"].HandleRoutes).Methods(http.MethodDelete).Name("deleteArtistsss")
	ar.HandleFunc("/", controllers["artists"].HandleRoutes).Methods(http.MethodGet).Name("getArtists")
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodGet).Name("getArtist")
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodPut).Name("updateArtist")
	ar.HandleFunc("/{id}", controllers["artists"].HandleRoutes).Methods(http.MethodDelete).Name("deleteArtist")

	return r
}
