package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DBConnection *sql.DB = nil

func getEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	username := getEnv("DB_USERNAME")
	password := getEnv("DB_PASSWORD")
	databse := getEnv("DB_DATABASE")

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", username, password, databse))

	if err != nil {
		panic(err.Error())
	}

	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	DBConnection = conn
}

func Close() {
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
