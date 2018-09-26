package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Schedule ok
type Schedule struct {
	Presentations []Presentation
}

// Event ok
type Event struct {
	Name        string
	Description string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
}

// Presentation ok
type Presentation struct {
	Event
	Speakers []string
	Topic    string
}

func newSchedule(url string) *Schedule {
	var sch Schedule
	xlp := newXMLParser(url)
	sch.Presentations = xlp.toPresentations()
	return &sch
}

func (sch *Schedule) handleScheduleAll(w http.ResponseWriter, req *http.Request) {
	log.Println("handleScheduleAll", req.RemoteAddr, req.Header)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err := enc.Encode(sch)
	if err != nil {
		log.Println("handleScheduleAll cannot encode schedule")
	}
}
