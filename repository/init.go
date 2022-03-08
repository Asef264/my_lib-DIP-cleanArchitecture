package repository

import (
	"database/sql"
	"fmt"
	"log"
	"my_lib/constants"

	_ "github.com/lib/pq"
)

//using postgrsql for db
//the main configuration of the database
// you should change the password and other different details in your project
const (
	user     = "postgres"
	password = "docker"
	host     = "localhost"
	port     = 5432
	dbname   = "postgres"
)

var DBP *sql.DB

func Connect() {
	dbConf := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user,
		password, host, port, dbname)

	db, err := sql.Open("postgres", dbConf)
	if err != nil {
		log.Println("error in connection function on sql.Open part")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("error in connection part and on pinging")
		panic(err)
	}
	fmt.Println("connection...")
	DBP = db
	_, err = db.Exec(constants.CREATE_BOOK_TABLE)
	if err != nil {
		log.Println("error on creating book table")
		panic(err)
	}
	log.Println("table : book status : created")
	_, err = db.Exec(constants.CREATE_BOOK_TABLE_CATEGORY)
	if err != nil {
		log.Println("error on creating category table")
		panic(err)
	}
	log.Println("category table created")
}
