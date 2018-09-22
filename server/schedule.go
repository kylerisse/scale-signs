package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type schedule []presentation

type event struct {
	name        string
	description string
	location    string
	startTime   time.Time
	endTime     time.Time
}

type presentation struct {
	event
	speakers []string
	audience string
	topic    string
}

func newSchedule(url string) *schedule {
	var sch schedule
	xlp := createXMLParser(url)
	sch = xlp.toPresentations()
	return &sch
}

func (sch *schedule) handleScheduleAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(sch)
}
