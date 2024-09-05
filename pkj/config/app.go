package config

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

func Connect() {
	database, err := sql.Open("mysql", "root:password@/postsdb")
	if err != nil {
		log.Fatal(err)
	}
	db = database
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
}
func GetDb() *sql.DB {
	return db
}
