package app

import (
	"log"

	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/belajar_gorm?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql2.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
