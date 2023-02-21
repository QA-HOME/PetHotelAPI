package postgre

import (
	sql2 "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"v1/src/errors"
	"v1/src/models/common"
	"v1/src/models/db"
)

func Connect() (*sql2.DB, *sql2.DB) {

	data := db.GetPostgresData()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.DBName)
	DB, err := sql2.Open("postgres", psqlInfo)
	errors.CheckErr(err)

	err = DB.Ping()
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("DB connection is succesfull!")
	}

	return DB, nil
}

func Close(db *sql2.DB) {
	err := db.Close()
	errors.CheckErr(err)
	if err == nil {
		log.Default().Print("DB Disconnected")
	}
}

func InsertUser(user common.User) error {

	PostgresDB, _ := Connect()
	defer Close(PostgresDB)

	log.Default().Print(user)

	result, err := PostgresDB.
		Exec("INSERT INTO Users(username, user_password, user_email, first_name, last_name, phone_number) VALUES($1, $2, $3, $4, $5, $6)",
			user.Username, user.UserPassword, user.UserEmail, user.FirstName, user.LastName, user.PhoneNumber)

	errors.CheckErr(err)

	rowsAffected, _ := result.RowsAffected()
	log.Default().Printf("Effected (%d) rows", rowsAffected)

	// ! LastInsertId is not supported by this driver
	lastID, err2 := result.LastInsertId()
	errors.CheckErr(err2)
	fmt.Println("\nLAST ID", lastID, err2)

	return err
}
