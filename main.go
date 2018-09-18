package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func stops(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func stopsPretty(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops_pretty.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func stations(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stations.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func stationsPretty(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stations_pretty.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func main() {

	/* port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	} */
	port := "8080"

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/stops", stops)
	http.HandleFunc("/stops_pretty", stopsPretty)
	http.HandleFunc("/stations", stations)
	http.HandleFunc("/stations_pretty", stationsPretty)
	http.ListenAndServe(":"+port, nil)

}
