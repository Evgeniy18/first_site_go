package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Port: 8080")
	http.ListenAndServe(":8080", nil)
}
