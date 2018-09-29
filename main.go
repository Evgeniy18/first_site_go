package main

import (
	"encoding/json"
	"html/template"
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

func stations(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/stations.json")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

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

func improvedServiceStatusSubway(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/improvedServiceStatusSubway.json")
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

func improvedServiceStatus(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("src/improvedServiceStatus.json")
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

func situations(w http.ResponseWriter, r *http.Request) {
	tmplPath := ""
	log.Print(r.URL.Path)
	switch r.URL.Path {
	case "/situations":
		tmplPath = "templates/dashboard.html"
	case "/situations.json":
		tmplPath = "templates/itemFeed.html"
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	type ViewData struct {
		Count      int
		Situations []Situation
	}

	myStruct := &ServiceStatusSubway{}
	fileNameOut := "src/improvedServiceStatusSubway.json"
	updateServiceStatusSubway(myStruct, fileNameOut)

	fin, err := ioutil.ReadFile("src/improvedServiceStatusSubway.json")
	if err != nil {
		log.Fatal(err)
	}

	var sits []Situation
	err = json.Unmarshal(fin, &sits)
	if err != nil {
		log.Fatal(err)
	}

	data := ViewData{
		Count:      len(sits),
		Situations: sits,
	}

	tmpl.Execute(w, data)
}

func serviceStatusRT(w http.ResponseWriter, r *http.Request) {
	tmplPath := ""
	log.Print(r.URL.Path)
	switch r.URL.Path {
	case "/service_status_rt":
		tmplPath = "templates/dashboardServiceStatus.html"
	case "/service_status_rt.json":
		tmplPath = "templates/itemServiceStatusFeed.html"
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	type ViewData struct {
		Count         int
		ServiceStatus ServiceStatus
	}

	myStruct := &Service{}
	fileNameOut := "src/improvedServiceStatus.json"
	updateServiceStatus(myStruct, fileNameOut)

	fin, err := ioutil.ReadFile("src/improvedServiceStatus.json")
	if err != nil {
		log.Fatal(err)
	}

	var servStatus ServiceStatus
	err = json.Unmarshal(fin, &servStatus)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(servStatus)
	data := ViewData{
		Count:         len(servStatus.Subway.Line),
		ServiceStatus: servStatus,
	}

	tmpl.Execute(w, data)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/stops", stops)
	http.HandleFunc("/stations", stations)
	http.HandleFunc("/calendar", calendar)
	http.HandleFunc("/calendar_dates", calendarDates)
	http.HandleFunc("/routes", routes)
	http.HandleFunc("/trips", trips)
	http.HandleFunc("/transfers", transfers)
	http.HandleFunc("/service_status_subway_1", serviceStatusSubway)
	http.HandleFunc("/service_status_subway_2", serviceStatusSubway2)
	http.HandleFunc("/service_status_subway_improved", improvedServiceStatusSubway)
	http.HandleFunc("/service_status_1", serviceStatus)
	http.HandleFunc("/service_status_2", serviceStatus2)
	http.HandleFunc("/service_status_improved", improvedServiceStatus)
	http.HandleFunc("/allequipments_1", allequipments)
	http.HandleFunc("/allequipments_2", allequipments2)
	http.HandleFunc("/nyct_ene_1", NYCEne)
	http.HandleFunc("/nyct_ene_2", NYCEne2)
	http.HandleFunc("/real_time_feed_1", RTFeed)

	http.HandleFunc("/situations", situations)
	http.HandleFunc("/situations.json", situations)
	http.HandleFunc("/service_status_rt", serviceStatusRT)
	http.HandleFunc("/service_status_rt.json", serviceStatusRT)

	http.ListenAndServe(":"+port, nil)

}
