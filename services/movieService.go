package services

import (
	"database/sql"
	"github.com/sh3lwan/webgo/config"
	"github.com/sh3lwan/webgo/models"
)

var db *sql.DB = config.DB()

func SelectMovies() []models.Movie {
	var movies []models.Movie = []models.Movie{}

	rows, err := db.Query("SELECT id, title FROM movies;")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var movie models.Movie
		rows.Scan(&movie.ID, &movie.Title)

		movies = append(movies, movie)
	}

	return movies
}
