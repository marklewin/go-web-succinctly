package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8999", http.FileServer(http.Dir("/var/www")))
}
