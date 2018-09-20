package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type sponsors struct {
	fileList []string
	lock     sync.Mutex
}

func newSponsors(imagePath string) *sponsors {
	var s sponsors
	sp := imagePath + "/sponsors"
	files, err := ioutil.ReadDir(sp)
	if err != nil {
		log.Fatal("cannot read " + sp)
	}
	for _, file := range files {
		s.fileList = append(s.fileList, file.Name())
	}
	return &s
}

func (s *sponsors) handleSponsorImageList(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.fileList)
}
