package main

import (
	"database/sql"
	"etee-api/config"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	fmt.Println("L'api est en cours d' execution sur le port 8000")

	config.DatabaseInit()

	defer config.Db().Close()

	router := InitializeRouter()

	http.ListenAndServe(":8000", router)
}
