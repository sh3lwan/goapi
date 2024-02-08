package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/sh3lwan/webgo/handlers"
)

func handleRoutes(r *mux.Router) {

	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies", controllers.AddMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.ShowMovie).Methods("GET")

}
