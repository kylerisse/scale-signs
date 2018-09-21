package server

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Nodes struct {
	all []Node `xml:"Nodes"`
}

// Node <node> from the upstream XML, pre cleanup
type Node struct {
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

func fetchXMLNodes(url string) Nodes {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("cannot get " + url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read XML response body " + url)
	}
	var nodes Nodes
	err = xml.Unmarshal(body, &nodes)
	if err != nil {
		log.Println("Unmarshal error")
		log.Println(err)
	}
	return nodes
}

func nodesToSchedulables(nodes Nodes) []presentation {
	var p []presentation
	for _, n := range nodes.all {
		p = append(p, nodeToPresentation(n))
	}
	return p
}

func nodeToPresentation(n Node) presentation {
	var p presentation
	p.name = n.Title
	p.description = removeHTMLTags(n.ShortAbstract)
	p.location = n.Room
	p.startTime = getStartTime(n.Day)
	p.endTime = getEndTime(n.Day)
	p.speakers = getSpeakers(n.Speakers)
	p.audience = "Everyone"
	p.topic = n.Topic
	return p
}

func removeHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

func getStartTime(ts string) time.Time {
	return time.Now()
}

func getEndTime(ts string) time.Time {
	return time.Now()
}

func getSpeakers(ss string) []string {
	var s []string
	return s
}
