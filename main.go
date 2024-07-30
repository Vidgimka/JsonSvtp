package main

import "fmt"

var geodatList []GeoData
var allRS []string
var loginConnect []LoginStat

func main() {
	fmt.Println("Парсинг json")
	geodatList = readFileData()
	allRS = uniqRSLists()
	loginConnect = createDB()
	fmt.Println(len(geodatList))
	fmt.Println(len(allRS))
	fmt.Println(loginConnect)
}
