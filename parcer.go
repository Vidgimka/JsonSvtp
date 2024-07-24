package main

import (
	"encoding/json"
	"log"
	"os"
)

func readFileData() {
	dat, err := os.ReadFile("./Svtp.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var geodat []GeoData

	err = json.Unmarshal(dat, &geodat)
	if err != nil {
		log.Fatal(err.Error())
	}
}
