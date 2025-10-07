package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Coffee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var coffees = []Coffee{
	{ID: "1", Name: "Espresso"},
	{ID: "2", Name: "Latte"},
	{ID: "3", Name: "Cappuccino"},
}

func getCoffees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func main() {
	http.HandleFunc("/coffees", getCoffees)
	log.Println("Coffee service listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
