package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo item structure
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos type is list of todo
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

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
