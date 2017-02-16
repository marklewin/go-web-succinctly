package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	log.Printf("Product ID:%v\n", productID)

	fileName := productID + ".html"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Printf("no such product")
		fileName = "invalid.html"
	}

	http.ServeFile(w, r, fileName)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", pageHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8999", nil)
}
