package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func main() {
	// replace "root" and "password" with your database login credentials
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/world")
	if err != nil {
		log.Println("Could not connect!")
	}
	database = db
	log.Println("Connected.")
}
