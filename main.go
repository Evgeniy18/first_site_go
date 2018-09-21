package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var files = [11]string{
	"calendar_dates",
	"calendar",
	"routes",
	"serviceStatus",
	"serviceStatus2",
	"serviceStatusSubway",
	"serviceStatusSubway2",
	"stations",
	"stops",
	"transfers",
	"trips",
}

func stops(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

/* func stopsPretty(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stops_pretty.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
} */

func stations(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stations.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

/* func stationsPretty(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stations_pretty.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
} */

func calendar(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/calendar.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func calendarDates(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/calendar_dates.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func routes(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/routes.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func serviceStatusSubway(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/ServiceStatusSubway.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func serviceStatusSubway2(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/ServiceStatusSubway2.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func serviceStatus(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/ServiceStatus.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func serviceStatus2(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/ServiceStatus2.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func allequipments(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/allequipments.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func allequipments2(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/allequipments2.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func NYCEne(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/nyct_ene.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func NYCEne2(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/nyct_ene2.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func trips(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/trips.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func transfers(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/transfers.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func RTFeed(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/feed_pretty.json")
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
	//http.HandleFunc("/stops_pretty", stopsPretty)
	http.HandleFunc("/stations", stations)
	//http.HandleFunc("/stations_pretty", stationsPretty)
	http.HandleFunc("/calendar", calendar)
	http.HandleFunc("/calendar_dates", calendarDates)
	http.HandleFunc("/routes", routes)
	http.HandleFunc("/trips", trips)
	http.HandleFunc("/transfers", transfers)
	http.HandleFunc("/service_status_subway_1", serviceStatusSubway)
	http.HandleFunc("/service_status_subway_2", serviceStatusSubway2)
	http.HandleFunc("/service_status_1", serviceStatus)
	http.HandleFunc("/service_status_2", serviceStatus2)
	http.HandleFunc("/allequipments_1", allequipments)
	http.HandleFunc("/allequipments_2", allequipments2)
	http.HandleFunc("/nyct_ene_1", NYCEne)
	http.HandleFunc("/nyct_ene_2", NYCEne2)
	http.HandleFunc("/real_time_feed_1", RTFeed)

	http.ListenAndServe(":"+port, nil)

}
