package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Decoders@123@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
