package server

import (
	"sync"
	"time"
)

type schedule struct {
	all  []presentation
	lock sync.Mutex
}

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
	nodes := fetchXMLNodes(url)
	sch.all = nodesToSchedulables(nodes)
	return &sch
}
