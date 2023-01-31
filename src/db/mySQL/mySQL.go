package mySQL

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"v1/src/errors"
)

const username = "root"
const password = "12345678"

func Connect() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:12345678@tcp(127.0.0.1:3306)/PetHotelAPI?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("MYSQL DB CONNECT")
	}
	return db, err
}

func Close(db *gorm.DB) error {

	mySqlDB, err := db.DB()
	errors.CheckErr(err)

	err = mySqlDB.Close()
	errors.CheckErr(err)

	if err == nil {
		log.Default().Print("MYSQL DB DISCONNECT")
	}
	return err
}
