package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Database() (*gorm.DB, error) {

	dsn := "user:password@tcp(localhost:3306)/halldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// if err = db.AutoMigrate(&Grocery{}); err != nil {
	//     log.Println(err)
	// }

	return db, err

}
