package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	InitDB("items.db")
	http.HandleFunc("/items", itemsHandler)
	log.Println("CRUD API server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func InitDB(s string) {
	panic("unimplemented")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		items, err := GetAllItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to fetch items"))
			return
		}
		json.NewEncoder(w).Encode(items)
	case http.MethodPost:
		var it Item
		if err := json.NewDecoder(r.Body).Decode(&it); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := InsertItem(it.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to insert item"))
			return
		}
		intID, ok := id.(int)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to convert id to int"))
			return
		}
		it.ID = intID
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(it)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func InsertItem(s string) (any, any) {
	panic("unimplemented")
}

func GetAllItems() (any, any) {
	panic("unimplemented")
}