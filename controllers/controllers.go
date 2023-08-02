package controllers

import (
	"encoding/json"
	"example/movieBackend/initializer"
	models "example/movieBackend/models"
	"fmt"
	"log"
	"time"

	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getting all movie list
func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movies []models.Movie
	initializer.Database.Find(&movies)

	json.NewEncoder(w).Encode(movies)
}

// delete any movie from database
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie models.Movie
	initializer.Database.First(&movie, params["id"])

	initializer.Database.Delete(&movie)
	var movies []models.Movie
	initializer.Database.Find(&movies)
	json.NewEncoder(w).Encode(movies)

}

// get any particular movie detail
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movie models.Movie
	initializer.Database.First(&movie, params["id"])

	json.NewEncoder(w).Encode(movie)
}

// create a movie record  to database
func CreateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie

	json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movie.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	movie.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	initializer.Database.Create(&movie)
	json.NewEncoder(w).Encode(movie)
}

// update any movie  record from database
func UpdateMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movie models.Movie
	initializer.Database.First(&movie, params["id"])

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&movie); err != nil {
		log.Fatal(err.Error())
		return
	}
	defer r.Body.Close()

	if err := initializer.Database.Save(&movie).Error; err != nil {
		log.Fatal(err.Error())
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// delete all movie  records from database
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	initializer.Database.Delete(&movie)
	json.NewEncoder(w).Encode(movie)

}

func IsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := getMoviewId(params["title"], params["release_date"])

	var movie models.Movie
	// initializer.Database.First(&movie, params["title"])
	initializer.Database.First(&movie, id)

	if movie.IsWatched {
		fmt.Fprintln(w, "current movie is watched earlier")
	} else {
		fmt.Fprintln(w, "current movie is not watched yet")
	}

	json.NewEncoder(w).Encode(movie)
}

func getMoviewId(title string, releaseDate string) string {

	var movie models.Movie

	initializer.Database.Where(map[string]interface{}{"title": title, "release_date": releaseDate}).First(&movie)

	return movie.ID
}
