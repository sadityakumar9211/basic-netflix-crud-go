package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sadityakumar9211/mongoapi/model"
	mymongo "github.com/sadityakumar9211/mongoapi/mongo"
)

// Actual controllers -  file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	movies := mymongo.GetAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var newMovie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		log.Fatal(err)
	}
	mymongo.InsertOneMovie(newMovie)
	fmt.Println("Inserted a movie")
	json.NewEncoder(w).Encode(newMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	mymongo.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	mymongo.DeleteOneMovie(params["id"])
	fmt.Println(fmt.Sprintf("Delete a movie with id: %v", params["id"]))
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteCount := mymongo.DeleteAllMovie()
	fmt.Println("The delete count is: ", deleteCount)
}
