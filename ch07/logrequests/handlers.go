package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the City Database!")
}

func CityList(w http.ResponseWriter, r *http.Request) {
	// Query the database
	jsonCities := dbCityList()

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCities)
}

func CityDisplay(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter with the city ID to search for
	vars := mux.Vars(r)
	cityId, _ := strconv.Atoi(vars["id"])

	// Query the database
	jsonCity := dbCityDisplay(cityId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCity)
}

func CityAdd(w http.ResponseWriter, r *http.Request) {
	var city City

	// Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &city); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422) // can't process!
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Write to the database
	addResult := dbCityAdd(city)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
}

func CityDelete(w http.ResponseWriter, r *http.Request) {

	// Get URL parameter with the city ID to delete
	vars := mux.Vars(r)
	cityId, _ := strconv.ParseInt(vars["id"], 10, 64)

	// Query the database
	deleteResult := dbCityDelete(cityId)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deleteResult)
}
