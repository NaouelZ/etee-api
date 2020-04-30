package main

import (
	"database/sql"
	"etee-api/config"
	"etee-api/models"
	"fmt"
	"net/http"

	_ "github.com/NaouelZ/etee-api/config"
	"github.com/gorilla/mux"

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

func InitializeRouter() *mux.Router {
	/* TO DO */

	router := mux.NewRouter()

	//Tickets routes
	router.HandleFunc("/tickets", models.GetTickets).Methods("GET")
	router.HandleFunc("/tickets", models.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", models.GetTicket).Methods("GET")
	router.HandleFunc("/tickets/{id}", models.UpdateTicket).Methods("PUT")
	router.HandleFunc("/tickets/{id}", models.DeleteTicket).Methods("DELETE")

	//Categories routes
	router.HandleFunc("/categories", models.GetCategories).Methods("GET")
	router.HandleFunc("/categories", models.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", models.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", models.DeleteCategory).Methods("DELETE")

	//Shops routes
	router.HandleFunc("/shops", models.GetShops).Methods("GET")
	router.HandleFunc("/shops", models.CreateShop).Methods("POST")
	router.HandleFunc("/shops/{id}", models.GetShop).Methods("GET")
	router.HandleFunc("/shops/{id}", models.DeleteShop).Methods("DELETE")

	//Shops Addresses routes
	router.HandleFunc("/address", models.GetAddresses).Methods("GET")
	router.HandleFunc("/address", models.CreateAddress).Methods("POST")
	router.HandleFunc("/address/{id}", models.GetAddress).Methods("GET")
	router.HandleFunc("/address/{id}", models.DeleteAddress).Methods("DELETE")

	return router
}
