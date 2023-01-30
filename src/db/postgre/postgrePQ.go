package postgre

import (
	sql2 "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"v1/src/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "PetHotelAPI"
)

func Connect() *sql2.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err := sql2.Open("postgres", psqlInfo)
	errors.CheckErr(err)

	err = DB.Ping()
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("DB connection is succesfull!")
	}

	return DB
}

func Close(db *sql2.DB) {
	err := db.Close()
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("DB Disconnected")
	}
}
