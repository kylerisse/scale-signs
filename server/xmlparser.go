package server

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type xmlParser struct {
	url   string
	nodes []Node
}

// Nodes <nodes> top level document
type Nodes struct {
	Nodes []Node `xml:"node"`
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

func newXMLParser(url string) *xmlParser {
	var xp xmlParser
	xp.url = url
	xp.nodes = xp.fetchXMLNodes()
	return &xp
}

func (xp *xmlParser) toPresentations() []Presentation {
	var ps []Presentation
	for i := 0; i < len(xp.nodes); i++ {
		ps = append(ps, xp.nodes[i].toPresentation())
	}
	return ps

}

func (n *Node) toPresentation() Presentation {
	var p Presentation
	p.Name = cleanupNewlinesAndSpaces(n.Title)
	p.Description = extractDescription(n.ShortAbstract)
	p.Location = cleanupNewlinesAndSpaces(n.Room)
	p.StartTime = extractStartTime(n.Time)
	p.EndTime = extractEndTime(n.Time)
	p.Speakers = extractSpeakers(n.Speakers)
	p.Topic = n.Topic
	return p
}

func (xp *xmlParser) fetchXMLNodes() []Node {
	resp, err := http.Get(xp.url)
	if err != nil {
		log.Println("cannot get " + xp.url)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("unable to close response body while fetching XML")
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read XML response body " + xp.url)
	}
	var xmlnodes Nodes
	err = xml.Unmarshal(body, &xmlnodes)
	if err != nil {
		log.Println("Unmarshal error")
		log.Println(err)
	}
	return xmlnodes.Nodes
}

func cleanupNewlinesAndSpaces(s string) string {
	rs := strings.TrimPrefix(s, "\n")
	rs = strings.TrimSuffix(rs, "\n")
	rs = strings.Replace(rs, "\n", " ", -1)
	rs = strings.Join(strings.Fields(rs), " ")
	rs = strings.TrimPrefix(rs, " ")
	rs = strings.TrimSuffix(rs, " ")
	return rs
}

func removeHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

func extractDescription(s string) string {
	return cleanupNewlinesAndSpaces(removeHTMLTags(s))
}

func extractTimeByIndex(s string, i int) time.Time {
	rs := strings.Split(s, " ")
	rt := strings.Split(rs[i], "\"")[1]
	t, err := time.Parse(time.RFC3339, rt)
	if err != nil {
		log.Println("error parsing timestring" + rt)
	}
	return t
}

func extractStartTime(s string) time.Time {
	return extractTimeByIndex(s, 5)
}

func extractEndTime(s string) time.Time {
	return extractTimeByIndex(s, 11)
}

func extractSpeakers(speakers string) []string {
	var rs []string
	for _, s := range strings.Split(speakers, ",") {
		rs = append(rs, cleanupNewlinesAndSpaces(s))
	}
	return rs
}
