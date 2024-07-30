package main

import "fmt"

func loginStatFill(login string) LoginStat { // наполниение данных о коннекте логина
	var l = LoginStat{
		Login:        login,
		TotalConnect: 0,
	}
	// считаем суммарно время коннекта
	for _, v := range geodatList {
		if login == v.Login {
			//l.SumConnecTime = l.SumConnecTime + v.Connect_time
			l.TotalConnect++
		}

	}

	return l
}

func createDB() []LoginStat {
	var db []LoginStat
	for _, v := range geodatList {
		db = append(db, loginStatFill(v.Login))
	}
	fmt.Println("Время коннекта посчитанно")
	return db
}
