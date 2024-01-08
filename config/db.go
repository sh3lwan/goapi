package config

import (
	"database/sql"
    "time"
	_ "github.com/go-sql-driver/mysql"
)

var DBConnection *sql.DB = nil

func Init() {
	conn, err := sql.Open("mysql", "root:@/webgo")

	if err != nil {
		panic(err.Error())
	}

	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
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
