package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaouelZ/etee-api/config"
	"github.com/gorilla/mux"
)

type Ticket struct {
	ID            int64   `json:"id"`
	PurchaseDate  string  `json:"purchase_date"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Name          string  `json:"name"`
	Commantary    string  `json:"commentary"`
	Pinned        bool    `json:"pinned"`
	UserID        int64   `json:"user_id"`
	ShopID        int64   `json:"shop_id"`
	CategoryID    int64   `json:"category_id"`
}

var db *sql.DB
var err error

//TO DO
// Distinguer les requettes HTTP ( mettre dans controller ) des actions a effectué

func GetTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tickets []Ticket

	result, err := config.Db().Query("SELECT * from tickets")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var ticket Ticket
		err := result.Scan(&ticket.ID, &ticket.PurchaseDate, &ticket.Amount, &ticket.PaymentMethod, &ticket.Name, &ticket.Commantary, &ticket.Pinned, &ticket.UserID, &ticket.ShopID, &ticket.CategoryID)
		if err != nil {
			panic(err.Error())
		}
		tickets = append(tickets, ticket)
	}

	json.NewEncoder(w).Encode(tickets)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var t Ticket

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := config.Db().Prepare("INSERT INTO tickets(purchase_date, amount, payment_method, name, commentary, pinned, user_id, shop_id, category_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(t.PurchaseDate, t.Amount, t.PaymentMethod, t.Name, t.Commantary, t.Pinned, t.UserID, t.ShopID, t.CategoryID)

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New ticket was created")
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := config.Db().Query("SELECT * FROM tickets WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var ticket Ticket
	for result.Next() {
		err := result.Scan(&ticket.ID, &ticket.PurchaseDate, &ticket.Amount, &ticket.PaymentMethod, &ticket.Name, &ticket.Commantary, &ticket.Pinned, &ticket.UserID, &ticket.ShopID, &ticket.CategoryID)

		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(ticket)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := config.Db().Prepare("UPDATE tickets SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	var t Ticket

	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = stmt.Exec(t.PurchaseDate, t.Amount, t.PaymentMethod, t.Name, t.Commantary, t.Pinned, t.UserID, t.ShopID, t.CategoryID, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Ticket with ID = %s was updated", params["id"])
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := config.Db().Prepare("DELETE FROM tickets WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Ticket with ID = %s was deleted", params["id"])
}
