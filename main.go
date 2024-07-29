package main

import "fmt"

var geodatList []GeoData
var allRS []string

func main() {
	fmt.Println("Парсинг json")
	geodatList = readFileData()
	allRS = uniqRSLists()
	fmt.Println(len(geodatList))
	fmt.Println(len(allRS))

}
