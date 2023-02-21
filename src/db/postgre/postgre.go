package postgre

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"v1/src/errors"
	db2 "v1/src/models/db"
)

func ConnectDB() (*gorm.DB, error) {

	pData := db2.GetPostgresData()
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		pData.Host, pData.User, pData.Password, pData.DBName, pData.Port, pData.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("POSTGRE DB IS CONNECTED")
	}
	return db, err
}

func CloseDB(db *gorm.DB) error {
	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("POSTGRE DB IS DISCONNECTED")
	}
	return err
}
