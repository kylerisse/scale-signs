package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type sponsors []string

func newSponsors(imagePath string) *sponsors {
	var s sponsors
	sp := imagePath + "/sponsors"
	files, err := ioutil.ReadDir(sp)
	if err != nil {
		log.Println("cannot read " + sp)
	}
	for _, file := range files {
		s = append(s, file.Name())
	}
	return &s
}

func (s *sponsors) handleSponsorImageList(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s)
}
