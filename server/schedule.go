package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type schedule []schedulable

func InitSchedule(s string) *schedule {
	var sch schedule
	resp, err := http.Get(s)
	if err != nil {
		log.Fatal("cannot get " + s)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("cannot read XML response body " + s)
	}
	var rawNodes RawNodes
	err = xml.Unmarshal(body, &rawNodes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", rawNodes)
	// rawnodetopres
	return &sch
}
