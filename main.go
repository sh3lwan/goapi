package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sh3lwan/webgo/config"
)

var db *sql.DB = config.DB()

type Data struct {
}

func parseTemplate(layout string) *template.Template {
	return template.Must(template.ParseFiles(layout))
}

func handleTemplate(layout string) func(w http.ResponseWriter, r *http.Request) {
	tmpl := parseTemplate(layout)
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{}
		tmpl.Execute(w, data)
	}
}

func GetEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	r := mux.NewRouter()

	handleRoutes(r)

	defer config.Close()
	log.Fatal(http.ListenAndServe(":8080", r))
}
