package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Name        string
	CountryCode string
	Population  uint32
}

var database *sql.DB

func main() {
	// replace "root" and "password" with your database login credentials
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/world")
	if err != nil {
		log.Println("Could not connect!")
	}
	database = db
	log.Println("Connected.")

	http.HandleFunc("/", showCity)
	http.ListenAndServe(":8999", nil)
}

func showCity(w http.ResponseWriter, r *http.Request) {
	city := City{}
	queryParam := "%" + r.URL.Path[1:] + "%"
	rows, err := database.Query("SELECT Name, CountryCode, Population FROM city WHERE Name LIKE ?", queryParam)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&city.Name, &city.CountryCode, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s (%s), Population: %d \n", city.Name, city.CountryCode, city.Population)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
