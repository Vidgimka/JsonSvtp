package main

import "fmt"

var geodatList []GeoData
var allRS []string
var loginConnect []LoginStat
var StationConnect []CalcRSstat

func main() {
	fmt.Println("Парсинг json")
	geodatList = readFileData()
	allRS = uniqRSLists()
	loginConnect = createDB()
	StationConnect = createDBRS()
	fmt.Println(len(geodatList))
	fmt.Println(len(allRS))
	fmt.Println(StationConnect[0])
}
