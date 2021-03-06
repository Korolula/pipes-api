package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func connectDB(connString string) *sql.DB {
	result, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func dbIsDown() bool {
	if _, err := db.Exec("SELECT 1"); err != nil {
		return true
	}
	return false
}
