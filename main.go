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
	Name      string
	Completed bool
	Due       time.Time
}

// Todos type is list of todo
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
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
	json.NewEncoder(w).Encode(todos)
}
