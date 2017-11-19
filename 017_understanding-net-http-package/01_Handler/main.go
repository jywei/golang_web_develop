package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

// hotdog is also a value to type hadler
// hotdog implement Habdler interface now

func main() {

}
