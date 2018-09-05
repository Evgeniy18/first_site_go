package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
	w.Write([]byte("!!!"))
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	fmt.Println("starting server at :" + port)
	http.ListenAndServe(":"+port, nil)
}
