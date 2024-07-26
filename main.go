package main

import "fmt"

var geodatList []GeoData

func main() {
	fmt.Println("Парсинг json")
	//fmt.Println(readFileData()) просто выводился джейсон в терминал
	geodatList = readFileData()

	//fmt.Println(geodatList[2]) выводился конкернтый объект джейсона
	fmt.Println(geodatList[2].Station_distance + 2) // один из элементов для последующих действий
}
