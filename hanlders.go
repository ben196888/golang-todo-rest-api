package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Index to response health
func Index(w http.ResponseWriter, r *http.Request) {
	msg := jsonMsg{200, "Welcome!"}
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

// TodoIndex to get all todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
