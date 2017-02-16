package main

import (
	"log"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "John Smith"}
	t, _ := template.ParseFiles("hello.html")

	t.Execute(w, p)
}
