package database

import (
	"JSONSVTP/model"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=SvtpPrin port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//gorm.Open("postgres", "user=postgres password=postgres dbname=SvtpPrin sslmode=disable TimeZone=Asia/Shanghai")

	if err != nil {
		fmt.Println("не подключилось к БД")
	}

	db.AutoMigrate(&model.SvtpPrin{})
	return db
}

func GetDB() *gorm.DB { // проверяем подкюченеие к базе данных
	if dbase == nil {
		dbase = Init() // проиничиализировали бд через Init и присвоили в переменную dbase, так как инициилизация return db
		var sleep = time.Duration(1)
		for dbase == nil { // ждем 3 секунды если не подключается
			sleep = sleep * 3
			fmt.Printf("База данных не доступна. Пожождите %d секунды.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init() // еще раз кладем базу данных в переменную dbase
		}
	}
	return dbase
}
