package config

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DBConnection *sql.DB = nil

func Init() {
	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}
	DBConnection = conn
}

func CloseConnection() {
	err := DBConnection.Close()
	if err != nil {
		panic(err)
	}
}

func DB() *sql.DB {
	if DBConnection == nil {
		Init()
	}

	return DBConnection
}
