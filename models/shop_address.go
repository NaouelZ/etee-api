package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"etee-api/config"

	_ "github.com/NaouelZ/etee-api/config"

	"github.com/gorilla/mux"
)

type ShopAddress struct {
	ID      int64  `json:"id"`
	Number  int64  `json:"number"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []ShopAddress
	result, err := config.Db().Query("SELECT * from address")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var s ShopAddress
		err := result.Scan(&s.ID, &s.Number, &s.Street, &s.City, &s.Country, &s.ZipCode)
		if err != nil {
			panic(err.Error())
		}
		shops = append(shops, s)
	}

	json.NewEncoder(w).Encode(shops)
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := config.Db().Query("SELECT * FROM address WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var s ShopAddress
	for result.Next() {
		err := result.Scan(&s.ID, &s.Number, &s.Street, &s.City, &s.Country, &s.ZipCode)

		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(s)
}

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var s ShopAddress

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := config.Db().Prepare("INSERT INTO address(number, street, city, country, zip_code) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(s.Number, s.Street, s.City, s.Country, s.ZipCode)

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Address has been add")
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := config.Db().Prepare("DELETE FROM address WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Address with ID = %s was deleted", params["id"])
}
