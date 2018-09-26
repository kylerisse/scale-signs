package main

import (
	"./server"
)

func main() {
	listenon := "127.0.0.1:8080"
	imagePath := "./images"
	xmlURL := "http://www.socallinuxexpo.org/scale/16x/sign.xml"

	server.StartServer(imagePath, xmlURL, listenon)
}
