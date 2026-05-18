package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// create handler functions
func getMovies(w http.ResponseWriter, r *http.Request) {

	// specify the value of the Header() field, Content-Type. This will interprete to the client the type of the response
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu := mux.Vars(r)
	for _, item := range movies {
		if item.ID == mu["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

// Use client request to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	// create a variable of type movie struct
	var mov Movie
	// Decode the client's json request body and store them in the variable
	_ = json.NewDecoder(r.Body).Decode(&mov)
	// use rand to generate a random number for the movie id
	mov.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, mov)
	json.NewEncoder(w).Encode(mov)

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParameters := mux.Vars(r)
	reqId := reqParameters["id"]
	for i, item := range movies {
		if reqId == item.ID {
			// delete the movie with the specified id
			movies = append(movies[:i], movies[i+1:]...)

			var mov Movie
			// decode data for the new movie
			_ = json.NewDecoder(r.Body).Decode(&mov)
			mov.ID = reqId
			// replace deleted movie with new one
			movies = append(movies, mov)
			json.NewEncoder(w).Encode(mov)
			return

		}
	}

}

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "555421", Title: "movie 2", Director: &Director{Firstname: "Kenedy", Lastname: "Rock"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Println("Starting server at port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

}
