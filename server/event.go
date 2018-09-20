package server

import (
	"time"
)

type schedulable interface {
	times() (time.Time, time.Time)
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

func (e *event) times() (time.Time, time.Time) {
	return e.startTime, e.endTime
}

func (p *presentation) times() (time.Time, time.Time) {
	return p.startTime, p.endTime
}
