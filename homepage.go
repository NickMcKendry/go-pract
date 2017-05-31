package main

import (
	"fmt"
	"net/http"
)

//comment

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {

	fmt.Println("Hello World")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
