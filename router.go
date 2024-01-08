package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sh3lwan/webgo/controllers"
)

func HandleRequests() {

	r := mux.NewRouter()
    

	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies", controllers.AddMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.ShowMovie).Methods("GET")

	r.Headers("Content-Type", "application/json")

	log.Fatal(http.ListenAndServe(":8080", r))

}
