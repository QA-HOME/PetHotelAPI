package mySQL

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"v1/src/errors"
	db2 "v1/src/models/db"
)

func Connect() (*gorm.DB, error) {

	dbData := db2.GetMySQLProperties()

	db, err := gorm.Open(mysql.Open(dbData.Dsn), &gorm.Config{})
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
