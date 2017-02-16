package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// Connect to the "world" database
func dbConnect() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/world")
	if err != nil {
		log.Println("Could not connect!")
	}
	database = db
	log.Println("Connected.")
}

// Find all cities and return as JSON
func dbCityList() []byte {
	var cities Cities
	var city City

	cityResults, err := database.Query("SELECT * FROM city")
	if err != nil {
		log.Fatal(err)
	}
	defer cityResults.Close()

	for cityResults.Next() {
		cityResults.Scan(&city.Id, &city.Name, &city.CountryCode, &city.District, &city.Population)
		cities = append(cities, city)
	}

	jsonCities, err := json.Marshal(cities)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCities
}

// Find a single city based on ID and return as JSON
func dbCityDisplay(id int) []byte {
	var city City

	err := database.QueryRow("SELECT * FROM city WHERE ID=?", id).Scan(&city.Id, &city.Name, &city.CountryCode, &city.District, &city.Population)
	if err != nil {
		log.Fatal(err)
	}

	jsonCity, err := json.Marshal(city)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCity
}

// Create a new city based on the information supplied
func dbCityAdd(city City) []byte {

	var addResult DBUpdate

	// Create prepared statement
	stmt, err := database.Prepare("INSERT INTO City(Name, CountryCode, District, Population) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(city.Name, city.CountryCode, city.District, city.Population)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	addResult.Id = lastId
	addResult.Affected = rowCnt

	// Convert to JSON and return
	newCity, err := json.Marshal(addResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return newCity
}

// Delete the city with the supplied ID
func dbCityDelete(id int64) []byte {
	var deleteResult DBUpdate

	// Create prepared statement
	stmt, err := database.Prepare("DELETE FROM City WHERE ID=?")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	deleteResult.Id = id
	deleteResult.Affected = rowCnt

	// Convert to JSON and return
	deletedCity, err := json.Marshal(deleteResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return deletedCity
}
