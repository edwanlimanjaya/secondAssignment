package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp")

	if err != nil {
		log.Fatal(err)
		fmt.Println("Could not connect to database")
	}

	return db
}
