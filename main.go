package main

import (
	"log"
	"net/http"
)

func main() {
	sp := initSponsors()
	//sch := initSchedule("http://www.socallinuxexpo.org/scale/16x/sign.xml")
	initSchedule("http://www.socallinuxexpo.org/scale/16x/sign.xml")
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/loadSponsorImages.js", sp.handleSponsorImages)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
