package main

import (
	"github.com/gorilla/mux"
	"github.com/sh3lwan/webgo/handlers"
)

func defineRoutes(r *mux.Router) {
    r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies", handlers.AddMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", handlers.ShowMovie).Methods("GET")
}
