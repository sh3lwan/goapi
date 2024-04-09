package config

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func parseTemplate(layout string) *template.Template {
	return template.Must(template.ParseFiles(layout))
}

func HandleTemplate(w http.ResponseWriter, layout string, data map[string]any) {
	tmpl := parseTemplate(layout)
	tmpl.Execute(w, data)
}
