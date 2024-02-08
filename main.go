package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sh3lwan/webgo/config"
)

var db *sql.DB = config.DB()

func main() {

	r := mux.NewRouter()

	handleRoutes(r)

    defer config.Close()

	log.Fatal(http.ListenAndServe(":8080", r))
}
