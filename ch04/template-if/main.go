package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Name        string
	CountryCode string
	Population  int
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
	var Cities = []City{}
	queryParam := "%" + r.URL.Path[1:] + "%"
	cities, err := database.Query("SELECT Name, CountryCode, Population FROM city WHERE Name LIKE ?", queryParam)
	if err != nil {
		log.Fatal(err)
	}
	defer cities.Close()

	for cities.Next() {
		theCity := City{}
		cities.Scan(&theCity.Name, &theCity.CountryCode, &theCity.Population)
		Cities = append(Cities, theCity)
	}

	t, _ := template.ParseFiles("results.html")

	t.Execute(w, Cities)
}
