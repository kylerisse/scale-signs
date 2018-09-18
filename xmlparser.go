package main

import (
	"regexp"
	"time"
)

// RawNodes <nodes> from the upstream XML
type RawNodes struct {
	Nodes []RawNode `xml:"node"`
}

// RawNode <node> from the upstream XML, pre cleanup
type RawNode struct {
	Title         string `xml:"Title"`
	Room          string `xml:"Room"`
	Day           string `xml:"Day"`
	Time          string `xml:"Time"`
	Speakers      string `xml:"Speakers"`
	SpeakerIDs    string `xml:"Speaker-IDs"`
	Topic         string `xml:"Topic"`
	ShortAbstract string `xml:"Short-abstract"`
	Photo         string `xml:"Photo"`
	Path          string `xml:"Path"`
}

/*
func rawNodeToPresentation(rn RawNode) presentation {
	start, end := extractTimes(rn.Time)
	// https://play.golang.org/p/mi-UxBSxq0U
	// https://play.golang.org/p/-9Hd7RQ3LwT
	return presentation{
		event{
			name:        rn.Title,
			description: removeHTMLTags(rn.ShortAbstract),
			location:    rn.Room,
		},
	}
}
*/

func removeHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

func extractTimes(ts string) (time.Time, time.Time) {
	return time.Now(), time.Now()
}
