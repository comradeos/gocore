package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>home</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>about</h1>")
}

func main() {
	fmt.Println("Simple Web Server")
	
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":89", nil)
}

