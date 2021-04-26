package driver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(host, user, password, dbname string) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database connection successful")

}
