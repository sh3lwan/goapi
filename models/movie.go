package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Auther *User  `json:"author"`
}

func NewMovie(title string) *Movie {
	return &Movie{
		Title: title,
	}
}
