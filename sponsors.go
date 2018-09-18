package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type sponsors []string

func initSponsors() *sponsors {
	var s sponsors
	filelist, err := ioutil.ReadDir("./images/sponsors")
	if err != nil {
		log.Fatal("can't read sponsor image dir")
	}
	for _, file := range filelist {
		s = append(s, file.Name())
	}
	return &s
}

func (s *sponsors) handleSponsorImages(w http.ResponseWriter, req *http.Request) {
	sponsorTemplate, err := template.ParseFiles("./templates/loadSponsorImages.gojs")
	if err != nil {
		log.Fatal("can't parse template file loadSponsorImages.gojs")
	}
	err = sponsorTemplate.Execute(w, s)
	if err != nil {
		log.Fatal("can't execute template file loadSponsorImages.gojs")
	}
}
