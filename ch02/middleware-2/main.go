package main

import (
	"net/http"
)

type AfterMiddleware struct {
	handler http.Handler
}

func (a *AfterMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	w.Write([]byte(" +++ Hello from middleware! +++ "))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" *** Hello from myHandler! *** "))
}

func main() {
	mid := &AfterMiddleware{http.HandlerFunc(myHandler)}

	println("Listening on port 8999")
	http.ListenAndServe(":8999", mid)
}
