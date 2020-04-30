package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/NaouelZ/etee-api/config"
)

func TestDatabaseInit(t *testing.T) {

	connection, err := sql.Open("mysql", "root")
	_, err = connection.Exec("CREATE DATABASE tickets_test")

	connection.Close()

	db, err = sql.Open("mysql", "root:@127.0.0.1:3306/tickets_test")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	config.CreateCarsTable()
}

func TestDatabaseDestroy(t *testing.T) {
	db.Close()

	connection, err := sql.Open("mysql", "root")
	_, err = connection.Exec("CREATE DATABASE tickets_test")

	if err != nil {
		log.Fatal(err)
	}
}
