package main

import (
	"encoding/json"
	"log"
	"os"

	"./server"
)

func main() {
	file, _ := os.Open("config.json")
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("unable to close filie config.json")
		}
	}()
	decoder := json.NewDecoder(file)
	conf := server.Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatalln("error:", err)
	}
	server.StartServer(&conf)
}
