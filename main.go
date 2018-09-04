package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
