package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type Siri struct {
	XMLName         xml.Name `xml:"Siri" json:"siri,omitempty"`
	Text            string   `xml:",chardata" json:"text,omitempty"`
	Ns2             string   `xml:"ns2,attr" json:"ns2,omitempty"`
	Xmlns           string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Ns4             string   `xml:"ns4,attr" json:"ns4,omitempty"`
	Ns3             string   `xml:"ns3,attr" json:"ns3,omitempty"`
	ServiceDelivery struct {
		Text              string `xml:",chardata" json:"text,omitempty"`
		ResponseTimestamp struct {
			Text string `xml:",chardata" json:"text,omitempty"`
		} `xml:"ResponseTimestamp" json:"responsetimestamp,omitempty"`
		SituationExchangeDelivery struct {
			Text              string `xml:",chardata" json:"text,omitempty"`
			ResponseTimestamp struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"ResponseTimestamp" json:"responsetimestamp,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Status" json:"status,omitempty"`
			Situations struct {
				Text               string `xml:",chardata" json:"text,omitempty"`
				PtSituationElement []struct {
					Text         string `xml:",chardata" json:"text,omitempty"`
					CreationTime struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"CreationTime" json:"creationtime,omitempty"`
					SituationNumber struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"SituationNumber" json:"situationnumber,omitempty"`
					PublicationWindow struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						StartTime struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"StartTime" json:"starttime,omitempty"`
						EndTime struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"EndTime" json:"endtime,omitempty"`
					} `xml:"PublicationWindow" json:"publicationwindow,omitempty"`
					Summary struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Summary" json:"summary,omitempty"`
					Description struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Description" json:"description,omitempty"`
					LongDescription struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"LongDescription" json:"longdescription,omitempty"`
					Planned struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"Planned" json:"planned,omitempty"`
					ReasonName struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"ReasonName" json:"reasonname,omitempty"`
					MessagePriority struct {
						Text string `xml:",chardata" json:"text,omitempty"`
					} `xml:"MessagePriority" json:"messagepriority,omitempty"`
					Source struct {
						Text       string `xml:",chardata" json:"text,omitempty"`
						SourceType struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"SourceType" json:"sourcetype,omitempty"`
					} `xml:"Source" json:"source,omitempty"`
					Affects struct {
						Text            string `xml:",chardata" json:"text,omitempty"`
						VehicleJourneys struct {
							Text                   string `xml:",chardata" json:"text,omitempty"`
							AffectedVehicleJourney []struct {
								Text    string `xml:",chardata" json:"text,omitempty"`
								LineRef struct {
									Text string `xml:",chardata" json:"text,omitempty"`
								} `xml:"LineRef" json:"lineref,omitempty"`
								DirectionRef struct {
									Text string `xml:",chardata" json:"text,omitempty"`
								} `xml:"DirectionRef" json:"directionref,omitempty"`
							} `xml:"AffectedVehicleJourney" json:"affectedvehiclejourney,omitempty"`
						} `xml:"VehicleJourneys" json:"vehiclejourneys,omitempty"`
					} `xml:"Affects" json:"affects,omitempty"`
					Consequences struct {
						Text        string `xml:",chardata" json:"text,omitempty"`
						Consequence []struct {
							Text      string `xml:",chardata" json:"text,omitempty"`
							Condition struct {
								Text string `xml:",chardata" json:"text,omitempty"`
							} `xml:"Condition" json:"condition,omitempty"`
							Severity struct {
								Text string `xml:",chardata" json:"text,omitempty"`
							} `xml:"Severity" json:"severity,omitempty"`
						} `xml:"Consequence" json:"consequence,omitempty"`
					} `xml:"Consequences" json:"consequences,omitempty"`
				} `xml:"PtSituationElement" json:"ptsituationelement,omitempty"`
			} `xml:"Situations" json:"situations,omitempty"`
		} `xml:"SituationExchangeDelivery" json:"situationexchangedelivery,omitempty"`
	} `xml:"ServiceDelivery" json:"servicedelivery,omitempty"`
}

type AffectedVehicleJourney struct {
	LineRef      string `json:"line_ref,omitempty"`
	DirectionRef string `json:"direction_ref,omitempty"`
}

type Consequence struct {
	Condition string `json:"condition,omitempty"`
	Severity  string `json:"severity,omitempty"`
}

type Situation struct {
	CreationTime           string                   `json:"creation_time,omitempty"`
	SituationNumber        string                   `json:"situation_number,omitempty"`
	StartTime              string                   `json:"start_time,omitempty"`
	EndTime                string                   `json:"end_time,omitempty"`
	Summary                string                   `json:"summary,omitempty"`
	Description            string                   `json:"description,omitempty"`
	LongDescription        template.HTML            `json:"long_description,omitempty"`
	Planned                string                   `json:"planned,omitempty"`
	ReasonName             string                   `json:"reason_name,omitempty"`
	MessagePriority        string                   `json:"message_priority,omitempty"`
	SourceType             string                   `json:"source_type,omitempty"`
	AffectedVehicleJourney []AffectedVehicleJourney `json:"affected_vehicle_journey,omitempty"`
	Consequences           []Consequence            `json:"consequences,omitempty"`
}

func situations(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Fatal(err)
	}

	type ViewData struct {
		Count      int
		Situations []Situation
	}

	myStruct := &Siri{}
	fileNameOut := "src/improvedServiceStatusSubway.json"
	saveXMLToJSONWithStruct(myStruct, fileNameOut)

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

func giveSituations(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/itemFeed.html")
	if err != nil {
		log.Fatal(err)
	}

	type ViewData struct {
		Count      int
		Situations []Situation
	}

	myStruct := &Siri{}
	fileNameOut := "src/improvedServiceStatusSubway.json"
	saveXMLToJSONWithStruct(myStruct, fileNameOut)

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

func saveXMLToJSONWithStruct(i *Siri, out string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://web.mta.info/status/ServiceStatusSubway.xml", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Print(resp.StatusCode)
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = xml.Unmarshal(body, i)
		if err != nil {
			log.Fatal(err)
		}

		//newService := improveServiceStatus(*i)
		newService := improveServiceStatusSubway(*i)
		json, err := json.MarshalIndent(newService, "", "	")
		if err != nil {
			log.Fatal(err)
		}

		fout, err := os.Create(out)
		if err != nil {
			log.Fatal(err)
		}
		defer fout.Close()
		fmt.Fprintf(bufio.NewWriter(fout), "%s", string(json))
		log.Print("File was refreshed.")
	}
}

func improveServiceStatusSubway(s Siri) []Situation {
	//situations := Situation{}
	serviceStatusSubway := []Situation{}
	for _, situation := range s.ServiceDelivery.SituationExchangeDelivery.Situations.PtSituationElement {
		newAffectedVehicleJourneys := []AffectedVehicleJourney{}
		for _, affectedVehicleJourney := range situation.Affects.VehicleJourneys.AffectedVehicleJourney {
			newAffectedVehicleJourney := AffectedVehicleJourney{
				DirectionRef: affectedVehicleJourney.DirectionRef.Text,
				LineRef:      affectedVehicleJourney.LineRef.Text,
			}
			newAffectedVehicleJourneys = append(newAffectedVehicleJourneys, newAffectedVehicleJourney)
		}
		newConsequences := []Consequence{}
		for _, consequence := range situation.Consequences.Consequence {
			newConsequence := Consequence{
				Condition: consequence.Condition.Text,
				Severity:  consequence.Severity.Text,
			}
			newConsequences = append(newConsequences, newConsequence)
		}
		newSituation := Situation{
			CreationTime:           situation.CreationTime.Text,
			SituationNumber:        situation.SituationNumber.Text,
			StartTime:              situation.PublicationWindow.StartTime.Text,
			EndTime:                situation.PublicationWindow.EndTime.Text,
			Summary:                situation.Summary.Text,
			Description:            situation.Description.Text,
			LongDescription:        template.HTML(situation.LongDescription.Text),
			Planned:                situation.Planned.Text,
			ReasonName:             situation.ReasonName.Text,
			MessagePriority:        situation.MessagePriority.Text,
			SourceType:             situation.Source.SourceType.Text,
			AffectedVehicleJourney: newAffectedVehicleJourneys,
			Consequences:           newConsequences,
		}
		serviceStatusSubway = append(serviceStatusSubway, newSituation)
	}
	return serviceStatusSubway
}

/* func updateFile() {
	for {
		time.Sleep(1 * time.Minute)
		myStruct := &Siri{}
		fileNameOut := "src/improvedServiceStatusSubway.json"
		saveXMLToJSONWithStruct(myStruct, fileNameOut)
	}
} */

func main() {

	//go updateFile()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	//port := "8080"

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
	http.HandleFunc("/situations.json", giveSituations)

	http.ListenAndServe(":"+port, nil)

}
