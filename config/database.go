package config

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	db, err = sql.Open("mysql", "root:@/tickets")

	fmt.Printf("La base de donnée a été connectée ! ")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
