package routes

import (
	"go-movies-crud/controllers"

	"github.com/gorilla/mux"
)

func MoviesRoute(router *mux.Router) {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
}
