package controllers

import (
	"encoding/json"
	"go-movies-crud/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.MoviesData)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.MoviesData {
		if item.ID == params["id"] {
			models.MoviesData = append(models.MoviesData[:index], models.MoviesData[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.MoviesData)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range models.MoviesData {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	models.MoviesData = append(models.MoviesData, movie)
	json.NewEncoder(w).Encode(models.MoviesData)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	params := mux.Vars(r)
	for index, item := range models.MoviesData {
		if item.ID == params["id"] {
			models.MoviesData = append(models.MoviesData[:index], models.MoviesData[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			models.MoviesData = append(models.MoviesData, movie)
			break
		}
	}
	json.NewEncoder(w).Encode(models.MoviesData)
}
