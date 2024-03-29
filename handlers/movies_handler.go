package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sh3lwan/webgo/errors"
	"github.com/sh3lwan/webgo/models"
	"github.com/sh3lwan/webgo/repositories"
)

var movies []models.Movie = []models.Movie{}

var authors []models.User = []models.User{}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, _ := repositories.SelectMovies()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")

	//author := CreateAuthor(r.FormValue("author"))

	movie := &models.Movie{
		Title: title,
	}

	movie, _ = repositories.CreateMovie(movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	paramId := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(paramId, 10, 64)

	if err != nil {
		json.NewEncoder(w).Encode(errors.NotFound())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	movie, err := repositories.GetMovie(id)

	if err != nil {
		json.NewEncoder(w).Encode(errors.NotFound())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func FindMovie(Id int) *models.Movie {
	var selectedMovie *models.Movie
	for _, movie := range movies {
		if movie.ID == Id {
			selectedMovie = &movie
		}
	}

	return selectedMovie
}

func FindLastId(searchable string) int {
	var lastId int
	if searchable == "authors" {
		for _, author := range authors {
			if author.ID > lastId {
				lastId = author.ID
			}
		}
	} else {
		for _, movie := range movies {
			if movie.ID > lastId {
				lastId = movie.ID
			}
		}
	}

	return lastId + 1
}

func CreateAuthor(name string) *models.User {
	var createdAuthor models.User
	for _, author := range authors {
		if author.Name == name {
			createdAuthor = author
		}
	}

	if createdAuthor.ID == 0 {
		createdAuthor = models.User{ID: FindLastId("authors"), Name: name}
		authors = append(authors, createdAuthor)
	}

	return &createdAuthor

}
