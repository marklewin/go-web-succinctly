package main

import (
	"net/http"
	"strconv"
	"time"
)

func CheckLastVisit(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("lastvisit") //

	expiry := time.Now().AddDate(0, 0, 1)

	cookie := &http.Cookie{
		Name:    "lastvisit",
		Expires: expiry,
		Value:   strconv.FormatInt(time.Now().Unix(), 10),
	}

	http.SetCookie(w, cookie)

	if err != nil {
		w.Write([]byte("Welcome to the site!"))
	} else {
		lasttime, _ := strconv.ParseInt(c.Value, 10, 0)
		html := "Welcome back! You last visited at: "
		html = html + time.Unix(lasttime, 0).Format("15:04:05")
		w.Write([]byte(html))
	}
}

func main() {
	http.HandleFunc("/", CheckLastVisit)
	http.ListenAndServe(":8999", nil)
}
