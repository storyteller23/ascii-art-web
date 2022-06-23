package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art", PostAsciiArt)

	log.Println("Server is listening on http://localhost:8181...")
	http.ListenAndServe(":8181", nil)
}
