package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("keep-it-secret-keep-it-safe"))

func handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values
	session.Values["abc"] = "cba"
	session.Values[111] = 222
	// Save the session values
	session.Save(r, w)
}

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/", handler)
	http.ListenAndServe(":8999", nil)
}
