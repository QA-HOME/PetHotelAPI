package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"v1/src/errors"
)

func ConnectDB() (*gorm.DB, error) {

	var dsn = "host=localhost user=postgres password=postgres dbname=PetHotelAPI port=5432 sslmode=disable"
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
