package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Database() (*gorm.DB, error) {

	dsn := "avvsrohith:AVVSR@mysqluser2003@tcp(127.0.0.1:3306)/halldb?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}



	return db, err

}
