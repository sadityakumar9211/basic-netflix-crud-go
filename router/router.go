package router

import (
	"github.com/gorilla/mux"
	"github.com/sadityakumar9211/mongoapi/controller"
)

func Router() *mux.Router {
	var r *mux.Router = mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie/create", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/update/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/delete-all", controller.DeleteAllMovie).Methods("DELETE")
	r.HandleFunc("/api/movie/delete/{id}", controller.DeleteOneMovie).Methods("DELETE")
	return r
}
