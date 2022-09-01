package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "rafael:password@tcp(database:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	con = db
	fmt.Println("Established Connection")
	return db
}
