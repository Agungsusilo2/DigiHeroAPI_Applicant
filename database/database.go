package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:agung@tcp(localhost:3306)/belajar_golang"))
	if err != nil {
		panic(err)
	}

	return database

}
