package server

import (
	"log"
	"net/http"
)

// Server is the main server runtime
type Server struct {
	imgPath string
	sp      *sponsors
	sch     *schedule
}

// StartServer returns a pointer to a new Server
func StartServer(imagePath string) {
	var s Server
	s.imgPath = imagePath
	s.sp = newSponsors(s.imgPath)

	// static content routes
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir(s.imgPath))))

	// api routes
	http.HandleFunc("/api/sponsors/images", s.sp.handleSponsorImageList)

	// listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
