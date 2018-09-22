package server

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type xmlParser []XMLNode

// XMLNodes <nodes> top level document
type XMLNodes struct {
	All []XMLNode `xml:"node"`
}

// XMLNode <node> from the upstream XML, pre cleanup
type XMLNode struct {
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

func createXMLParser(url string) *xmlParser {
	var xp xmlParser
	//nodes := fetchXMLNodes(url)
	//fmt.Println(nodes)
	return &xp
}

func (xp *xmlParser) toPresentations() []presentation {
	var ps []presentation
	/*for i := 0; i < .length; i++ {
		ps.append(ps, xp[i].toPresentation())
	}*/
	return ps

}

func (n *XMLNode) toPresentation() presentation {
	var p presentation
	p.name = n.Title
	p.description = removeHTMLTags(n.ShortAbstract)
	p.location = n.Room
	p.startTime = extractStartTime(n.Day)
	p.endTime = extractEndTime(n.Day)
	p.speakers = extractSpeakers(n.Speakers)
	p.audience = "Everyone"
	p.topic = n.Topic
	return p
}

func fetchXMLNodes(url string) []XMLNode {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("cannot get " + url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//log.Println("fetch", url)
	//log.Println(resp.Status, resp.Header)
	if err != nil {
		log.Println("cannot read XML response body " + url)
	}
	var xmlnodes XMLNodes
	err = xml.Unmarshal(body, &xmlnodes)
	if err != nil {
		log.Println("Unmarshal error")
		log.Println(err)
	}
	return xmlnodes.All
}

func removeHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

func extractStartTime(ts string) time.Time {
	return time.Now()
}

func extractEndTime(ts string) time.Time {
	return time.Now()
}

func extractSpeakers(ss string) []string {
	var s []string
	return s
}
