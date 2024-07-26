package main

import (
	"encoding/json"
	"log"
	"os"
)

func readFileData() []GeoData {
	dat, err := os.ReadFile("./SvtpPrin.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var geodat []GeoData

	err = json.Unmarshal(dat, &geodat)
	if err != nil {
		log.Fatal(err.Error())
	}
	return geodat
}
