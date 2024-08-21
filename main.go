package main

import (
	"JSONSVTP/database"
	"fmt"
)

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
	// юзаем горм для создания и записи БД
	database.Init()
}
