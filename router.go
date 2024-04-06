package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/sh3lwan/webgo/handlers"
)


func handleRoutes(r *mux.Router) {
	r.HandleFunc("/", handleTemplate("layouts/index.html"))

	r.HandleFunc("/api/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies", controllers.AddMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controllers.ShowMovie).Methods("GET")

}
