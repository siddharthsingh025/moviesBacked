package main

import (
	controllers "example/movieBackend/controllers"
	"example/movieBackend/initializer"
	models "example/movieBackend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies []models.Movie

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.UpdateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovies).Methods("DELETE")
	r.HandleFunc("/movies", controllers.DeleteAllMovies).Methods("DELETE")
	r.HandleFunc("/movies/iswatched/", controllers.IsWatched).Methods("GET")

	fmt.Printf("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
