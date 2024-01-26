package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "7675", Title: "Kashmir Files", Director: &Director{FirstName: "Vivek", LastName: "Agnihotri"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7625", Title: "Hanuman", Director: &Director{FirstName: "Prashant", LastName: "Verma"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at")
	log.Fatal(http.ListenAndServe(":8000", r))
}
