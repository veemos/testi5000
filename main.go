package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandeFunc)
	http.ListenAndServe("localhost:8080", nil)
	http.ListenAndServe("localhost:8080", nil)
}

func helloHandeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
