package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/sh3lwan/webgo/config"
)

var db *sql.DB = config.DB()

type RenderComponent struct {
	templ.Component
}

func (c RenderComponent) RenderHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

}

func main() {
	r := mux.NewRouter()

	defineRoutes(r)

	fs := http.FileServer(http.Dir("./public/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	defer config.Close()

	log.Fatal(http.ListenAndServe(":8080", r))
}
