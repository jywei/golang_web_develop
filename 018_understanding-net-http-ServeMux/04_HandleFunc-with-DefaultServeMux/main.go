package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
