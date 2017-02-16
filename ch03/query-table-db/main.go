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

	db, err := sql.Open("mysql", "root:oracle@tcp(127.0.0.1:3306)/world")
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

	html := "<html><head><title>City Search</title></head><body><h1>Search for" + queryParam + "</h1><table border='1'><tr><th>City</th><th>Country Code</th><th>Population</th></tr>"

	for rows.Next() {
		err := rows.Scan(&city.Name, &city.CountryCode, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td></tr>", city.Name, city.CountryCode, city.Population)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	} else {
		html += "</table></body></html>"
		fmt.Fprintln(w, html)
	}
}
