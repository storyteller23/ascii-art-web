package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art", PostAsciiArt)
	port := ":" + os.Getenv("PORT")
	log.Println("Server is listening on http://127.0.0.1:8181...")
	http.ListenAndServe(port, nil)
}
