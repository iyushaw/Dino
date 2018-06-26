package main

import (
	"dino/dynowebportal"
	"encoding/json"
	"os"
	"log"
)

type configuration struct{
	ServerAddress string `json:"webserver"`
}

func main(){
	file,err := os.Open("config.json")
	if err!=nil{
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting Webserver on", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}