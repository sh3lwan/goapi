package services

import (
	"database/sql"
	"github.com/sh3lwan/webgo/config"
	. "github.com/sh3lwan/webgo/models"
)

var db *sql.DB = config.DB()

func SelectMovies() []Movie {
	var movies []Movie = []Movie{}

	rows, err := db.Query("SELECT id, title FROM movies;")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.ID, &movie.Title)

		movies = append(movies, movie)
	}

	return movies
}

func CreateMovie(movie Movie) Movie {

	statment, err := db.Prepare("INSERT INTO movies(title) VALUES (?);")

	if err != nil {
		panic(err.Error())
	}

	insertResult, err := statment.Exec(movie.Title)

	if err != nil {
		panic(err.Error())
	}

	lastInsertedId, err := insertResult.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	movie, _ = GetMovie(lastInsertedId)

	return movie
}

func GetMovie(movieId int64) (Movie, error) {
	fetchedMovie := Movie{}

	err := db.QueryRow("SELECT id, title FROM movies WHERE id = ?", movieId).Scan(&fetchedMovie.ID, &fetchedMovie.Title)

	return fetchedMovie, err
}
