package router

import (
	"github.com/Sunny955/mongoapi/controller"
	middleware "github.com/Sunny955/mongoapi/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.Logger)

	// Create a subgroup for /api/movies routes
	moviesSubrouter := router.PathPrefix("/api/movies").Subrouter()
	// moviesSubrouter.Use(middleware.Logger)
	moviesSubrouter.HandleFunc("", controller.GetMyAllMovies).Methods("GET")
	moviesSubrouter.HandleFunc("", controller.CreateMovie).Methods("POST")

	// router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	// router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")

	movieSubrouter := router.PathPrefix("/api/movie").Subrouter()
	// movieSubrouter.Use(middleware.Logger)
	movieSubrouter.HandleFunc("/{id}", controller.MarkedAsWatch).Methods("PUT")
	movieSubrouter.HandleFunc("/{id}", controller.DeleteAMovie).Methods("DELETE")

	// router.HandleFunc("/api/movie/{id}", controller.MarkedAsWatch).Methods("PUT")
	// router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/delete-all", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
