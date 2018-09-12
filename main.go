package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func stops(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(f)
}

func stopsPretty(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops_pretty.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(f)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/stops", stops)
	http.HandleFunc("/stops_pretty", stopsPretty)
	http.ListenAndServe(":"+port, nil)

}
