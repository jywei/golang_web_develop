package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type hotcat int

func (d hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
