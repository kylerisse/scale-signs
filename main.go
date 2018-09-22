package main

import (
	"./server"
)

func main() {
	port := ":8080"
	imagePath := "./images"
	xmlURL := "http://www.socallinuxexpo.org/scale/16x/sign.xml"

	server.StartServer(imagePath, xmlURL, port)
}
