package main

import (
	"fmt"
	"net/http"
	"time"

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

	// Get any previous flashes.
	if flashes := session.Flashes(); len(flashes) > 0 {
		// Do something with them
		for f := range flashes {
			fmt.Println(flashes[f])
		}
	} else {
		// Set a new flash.
		session.AddFlash("Flash! Ah-ah, savior of the universe! " + time.Now().String())
	}
	session.Save(r, w)
}

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/", handler)
	http.ListenAndServe(":8999", nil)
}
