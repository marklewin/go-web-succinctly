package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	dbConnect()

	log.Fatal(http.ListenAndServe(":8999", router))
}
