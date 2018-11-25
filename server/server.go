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

// StartServer starts the web server
func StartServer(c *Config) {
	var s Server

	log.Println("starting server")
	log.Println("opts: listenon=" + c.Listenon + " imagePath=" + c.ImagePath)
	log.Println("opts: xmlURL=" + c.XMLurl)

	s.sp = newSponsors(c.ImagePath)
	s.sch = newSchedule(c.XMLurl)
	s.rms = newRooms(s.sch)

	// static content routes
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir(c.ImagePath))))

	// api routes
	http.HandleFunc("/api/sponsors", s.sp.handleSponsorImageList)
	http.HandleFunc("/api/schedule", s.sch.handleScheduleAll)

	// listen
	log.Printf("listening on %s", c.Listenon)
	log.Fatal(http.ListenAndServe(c.Listenon, nil))
}
