package main

import (
	"bytes"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Name        string
	CountryCode string
	Population  int
}

func (c City) FormatPopulation(n int) string {

	sep := ','
	log.Println("hello")
	s := strconv.Itoa(n)

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {

		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++

		buff.WriteByte(s[i])
	}

	log.Println(buff.String())
	return buff.String()

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
