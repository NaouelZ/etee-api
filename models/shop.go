package models

import (
	"encoding/json"
	"etee-api/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Shop struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AddressID int64  `json:"address_id"`
}

func GetShops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []Shop
	result, err := config.Db().Query("SELECT * from shops")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var s Shop
		err := result.Scan(&s.ID, &s.Name, &s.AddressID)
		if err != nil {
			panic(err.Error())
		}
		shops = append(shops, s)
	}

	json.NewEncoder(w).Encode(shops)
}

func GetShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := config.Db().Query("SELECT * FROM shops WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var s Shop
	for result.Next() {
		err := result.Scan(&s.ID, &s.Name)

		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(s)
}

func CreateShop(w http.ResponseWriter, r *http.Request) {
	var s Shop

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := config.Db().Prepare("INSERT INTO shops(name, address_id) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(s.Name, s.AddressID)

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New shop was created")
}

func DeleteShop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := config.Db().Prepare("DELETE FROM shops WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Shop with ID = %s was deleted", params["id"])
}
