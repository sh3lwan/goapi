package repositories

import (
	"database/sql"

	"github.com/sh3lwan/webgo/config"
	"github.com/sh3lwan/webgo/models"
	. "github.com/sh3lwan/webgo/models"
)

var db *sql.DB = config.DB()

func SelectMovies() ([]Movie, error) {
	var movies []models.Movie = []models.Movie{}
	rows, err := db.Query("SELECT id, title FROM movies;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.ID, &movie.Title)
		movies = append(movies, movie)
	}

	return movies, nil
}

func CreateMovie(movie *Movie) (*Movie, error) {
	statment, err := db.Prepare("INSERT INTO movies(title) VALUES (?);")

	if err != nil {
		return nil, err
	}

	result, err := statment.Exec(movie.Title)

	if err != nil {
		return nil, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return GetMovie(lastInsertedId)
}

func GetMovie(movieId int64) (*Movie, error) {
	var movie Movie
	err := db.QueryRow("SELECT id, title FROM movies WHERE id = ?", movieId).Scan(&movie.ID, &movie.Title)

	return &movie, err
}
