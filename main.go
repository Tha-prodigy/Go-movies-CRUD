package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []movie 
func getMovies() {

}
func getMovie() {

}
func createMovie() {

}
func deleteMovie() {

}
func updateMovie() {
	
}

func main() {

	movies = append(movies, movie{ID: "1", Isbn: "438227", Title: "movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, movie{ID: "2", Isbn: "555421", Title: "movie 2", Director: &Director{Firstname: "Kenedy", Lastname: "Rock"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")

	fmt.Println("Starting server at port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

	
}
