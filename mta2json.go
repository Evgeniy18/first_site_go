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

//ServiceStatusSubway.xml
type ServiceStatusSubway struct {
	XMLName         xml.Name `xml:"Siri" json:"ServiceStatusSubway,omitempty"`
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

//improved ServiceStatusSubway struct
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

//ServiceStatus.txt
type Service struct {
	XMLName      xml.Name `xml:"service" json:"service,omitempty"`
	Text         string   `xml:",chardata" json:"text,omitempty"`
	Responsecode struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"responsecode" json:"responsecode,omitempty"`
	Timestamp struct {
		Text string `xml:",chardata" json:"text,omitempty"`
	} `xml:"timestamp" json:"timestamp,omitempty"`
	Subway struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Line []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"name" json:"name,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"status" json:"status,omitempty"`
			Text struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"text" json:"text,omitempty"`
			Date struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Date" json:"date,omitempty"`
			Time struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Time" json:"time,omitempty"`
		} `xml:"line" json:"line,omitempty"`
	} `xml:"subway" json:"subway,omitempty"`
	Bus struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Line []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"name" json:"name,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"status" json:"status,omitempty"`
			Text struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"text" json:"text,omitempty"`
			Date struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Date" json:"date,omitempty"`
			Time struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Time" json:"time,omitempty"`
		} `xml:"line" json:"line,omitempty"`
	} `xml:"bus" json:"bus,omitempty"`
	BT struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Line []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"name" json:"name,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"status" json:"status,omitempty"`
			Text struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"text" json:"text,omitempty"`
			Date struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Date" json:"date,omitempty"`
			Time struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Time" json:"time,omitempty"`
		} `xml:"line" json:"line,omitempty"`
	} `xml:"BT" json:"bt,omitempty"`
	LIRR struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Line []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"name" json:"name,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"status" json:"status,omitempty"`
			Text struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"text" json:"text,omitempty"`
			Date struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Date" json:"date,omitempty"`
			Time struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Time" json:"time,omitempty"`
		} `xml:"line" json:"line,omitempty"`
	} `xml:"LIRR" json:"lirr,omitempty"`
	MetroNorth struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Line []struct {
			Chardata string `xml:",chardata" json:"chardata,omitempty"`
			Name     struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"name" json:"name,omitempty"`
			Status struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"status" json:"status,omitempty"`
			Text struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"text" json:"text,omitempty"`
			Date struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Date" json:"date,omitempty"`
			Time struct {
				Text string `xml:",chardata" json:"text,omitempty"`
			} `xml:"Time" json:"time,omitempty"`
		} `xml:"line" json:"line,omitempty"`
	} `xml:"MetroNorth" json:"metronorth,omitempty"`
}

//improved ServiceStatus struct
type ServiceStatus struct {
	Subway Subway `json:"subway"`
	//	Bus        Bus        `json:"bus"`
	//	BT         BT         `json:"BT"`
	//	LIRR       LIRR       `json:"LIRR"`
	//	MetroNorth MetroNorth `json:"metro_north"`
}

type Subway struct {
	Line []Line `json:"line"`
}

/* type Bus struct {
	Line []Line `json:"line"`
}

type BT struct {
	Line []Line `json:"line"`
}

type LIRR struct {
	Line []Line `json:"line"`
}

type MetroNorth struct {
	Line []Line `json:"line"`
} */

type Line struct {
	Name   string        `json:"name,omitempty"`
	Status string        `json:"status,omitempty"`
	Text   template.HTML `json:"text,omitempty"`
	Date   string        `json:"date,omitempty"`
	Time   string        `xml:"Time" json:"time,omitempty"`
}

func updateServiceStatusSubway(i *ServiceStatusSubway, out string) {

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

		newService := improveServiceStatusSubway(*i)
		json, err := json.Marshal(newService)
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

func improveServiceStatusSubway(s ServiceStatusSubway) []Situation {
	situations := []Situation{}
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
		situations = append(situations, newSituation)
	}
	return situations
}

func updateServiceStatus(i *Service, out string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://web.mta.info/status/serviceStatus.txt", nil)
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
		log.Print("1111111")
		err = xml.Unmarshal(body, i)
		if err != nil {
			log.Fatal(err)
		}

		newService := improveServiceStatus(*i)
		json, err := json.Marshal(newService)
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

func improveServiceStatus(s Service) ServiceStatus {
	serviceStatus := &ServiceStatus{}
	for _, line := range s.Subway.Line {
		newLine := Line{
			Name:   line.Name.Text,
			Status: line.Status.Text,
			Text:   template.HTML(line.Text.Text),
			Date:   line.Date.Text,
			Time:   line.Time.Text,
		}
		serviceStatus.Subway.Line = append(serviceStatus.Subway.Line, newLine)
	}
	/* for _, line := range s.Bus.Line {
		newLine := Line{
			Name:   line.Name.Text,
			Status: line.Status.Text,
			Text:   line.Text.Text,
			Date:   line.Date.Text,
			Time:   line.Time.Text,
		}
		serviceStatus.Bus.Line = append(serviceStatus.Bus.Line, newLine)
	}
	for _, line := range s.BT.Line {
		newLine := Line{
			Name:   line.Name.Text,
			Status: line.Status.Text,
			Text:   line.Text.Text,
			Date:   line.Date.Text,
			Time:   line.Time.Text,
		}
		serviceStatus.BT.Line = append(serviceStatus.BT.Line, newLine)
	}
	for _, line := range s.LIRR.Line {
		newLine := Line{
			Name:   line.Name.Text,
			Status: line.Status.Text,
			Text:   line.Text.Text,
			Date:   line.Date.Text,
			Time:   line.Time.Text,
		}
		serviceStatus.LIRR.Line = append(serviceStatus.LIRR.Line, newLine)
	}
	for _, line := range s.MetroNorth.Line {
		newLine := Line{
			Name:   line.Name.Text,
			Status: line.Status.Text,
			Text:   line.Text.Text,
			Date:   line.Date.Text,
			Time:   line.Time.Text,
		}
		serviceStatus.MetroNorth.Line = append(serviceStatus.MetroNorth.Line, newLine)
	} */
	return *serviceStatus
}
