package controllers

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	// db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/db_name?parseTime=true&loc=Asia%2FJakarta")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_notflex?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
