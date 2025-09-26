package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func helloJenkinsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Jenkins!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/jenkins", helloJenkinsHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}


