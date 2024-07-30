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

// наполнение данных о коннекте станции
func RSStatFill(login string) CalcRSstat {
	var t = CalcRSstat{
		Login:         login,
		SumConnecTime: 0,
		MinConnecTime: 0,
		MaxConnecTime: 0,
	}
	// считаем суммарно время коннекта
	for _, v := range geodatList {
		if login == v.Login {
			t.SumConnecTime = t.SumConnecTime + v.Connect_time // сумируем время подключения
		}

	}
	return t
}
func createDBRS() []CalcRSstat {
	var dbrs []CalcRSstat
	for _, v := range geodatList {
		//dbrs = append(dbrs, RSStatFill(v))
		dbrs = append(dbrs, RSStatFill(v.Login))
	}
	fmt.Println("Время коннекта по станциям посчитанно")
	return dbrs
}
