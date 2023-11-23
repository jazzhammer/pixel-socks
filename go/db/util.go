package db

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB = nil

func StartDb() {
	dburl := "postgresql://postgres@localhost/pixel_socks?sslmode=disable"
	database, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
		fmt.Println(fmt.Sprintf("database connection: failed: %s", err))
	} else {
		fmt.Println("database connection: ok")
	}
	Db = database
}
