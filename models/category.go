package models

import (
	"encoding/json"
	"etee-api/config"
	"fmt"
	"net/http"

	_ "github.com/NaouelZ/etee-api/config"

	"github.com/gorilla/mux"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categories []Category
	result, err := config.Db().Query("SELECT * from categories")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var category Category
		err := result.Scan(&category.ID, &category.Name)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c Category

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := config.Db().Prepare("INSERT INTO categories(name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(c.Name)

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New category was created")
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := config.Db().Query("SELECT * FROM categories WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var c Category
	for result.Next() {
		err := result.Scan(&c.ID, &c.Name)

		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(c)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := config.Db().Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Category with ID = %s was deleted", params["id"])
}
