package main

import (
	"./server"
)

func main() {
	imagePath := "./images"
	//signURL := "http://www.socallinuxexpo.org/scale/16x/sign.xml"
	server.StartServer(imagePath)
}
