package main

import "fmt"

func loginStatFill(login string) LoginStat { // наполниение данных о коннекте логина
	var l = LoginStat{
		Login:         login,
		SumConnecTime: 0,
		MinConnecTime: 0,
		MaxConnecTime: 0,
	}
	// считаем суммарно время коннекта
	for _, v := range geodatList {
		if login == v.Login {
			l.SumConnecTime = l.SumConnecTime + v.Connect_time
		}

	}
	// считаем минимальное время коннекта

	// считаем максимальное время коннекта

	return l
}

func createDB() []LoginStat {
	var db []LoginStat
	for _, v := range allRS {
		db = append(db, loginStatFill(v))
	}
	fmt.Println("Время коннекта посчитанно")
	return db
}
