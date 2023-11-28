package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



type Movie struct {
	Title    string `json:"title"`
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Director *Director `json:"director"`
}


type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// * Get Movies Handler
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// * Get Movie Hanlder (id)
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, movie := range movies {
		if movie.ID == key {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// * Create Movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

// * Update Movie route Handler
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	for index, item := range movies {
		if item.ID == key {
			movies = append(movies[:index], movies[index+1:]...)
			var updated Movie
			_ = json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = key
			movies = append(movies, updated)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for index, item := range movies {
		if item.ID == key {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			break 
		}
	}
}

// * Added intial Movies
func addMovies() {
	movies = append(movies, Movie{ID: "1", Isbn: "321456", Title: "Movie One", Director: &Director{Firstname: "Arnold", Lastname: "Scrcemuller"}})
	movies = append(movies, Movie{ID: "2", Isbn: "321789", Title: "Movie Two", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
}

func main() {
	addMovies()
	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5050", router))
}