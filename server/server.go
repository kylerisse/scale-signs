package server

import (
	"log"
	"net/http"
)

// Server is the main server runtime
type Server struct {
	sp  *sponsors
	sch *Schedule
	rms *rooms
}

// StartServer returns a pointer to a new Server
func StartServer(imagePath string, xmlURL string, listenon string) {
	var s Server

	log.Println("starting server")
	log.Println("opts: listenon=" + listenon + " imagePath=" + imagePath)
	log.Println("opts: xmlURL=" + xmlURL)

	s.sp = newSponsors(imagePath)
	s.sch = newSchedule(xmlURL)
	s.rms = newRooms(s.sch)

	// static content routes
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir(imagePath))))

	// api routes
	http.HandleFunc("/api/sponsors", s.sp.handleSponsorImageList)
	http.HandleFunc("/api/schedule", s.sch.handleScheduleAll)

	// listen
	log.Printf("listening on %s", listenon)
	log.Fatal(http.ListenAndServe(listenon, nil))
}
