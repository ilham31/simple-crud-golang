package main

import (
	"fmt"
	"go-movies-crud/models"
	"go-movies-crud/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// init movies
	models.MoviesData = append(models.MoviesData, models.Movie{ID: "1", ISBN: "438227", Title: "Movie One", Director: &models.Director{Firstname: "John", Lastname: "Doe"}})
	models.MoviesData = append(models.MoviesData, models.Movie{ID: "2", ISBN: "454555", Title: "Movie Two", Director: &models.Director{Firstname: "Steve", Lastname: "Smith"}})

	// setup mux
	router := mux.NewRouter()
	routes.MoviesRoute(router)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
