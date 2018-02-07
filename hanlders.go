package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	var todoID int
	var err error
	if todoID, err = strconv.Atoi(params["todoID"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoID)
	// check not empty todo
	if todo.ID > 0 {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// return error when got empty todo
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonMsg{
		Code: http.StatusNotFound,
		Msg:  "Not Found",
	}); err != nil {
		panic(err)
	}
}
