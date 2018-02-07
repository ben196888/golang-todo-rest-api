package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index to response health
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// TodoIndex to get all todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow to show todo by todo id
func TodoShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["todoID"]
	if err := json.NewEncoder(w).Encode(todoID); err != nil {
		panic(err)
	}
}
