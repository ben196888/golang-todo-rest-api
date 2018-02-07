package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const kiloBytes int64 = 1024
const megaBytes int64 = 1024 * kiloBytes

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
	todoID, err := strconv.Atoi(params["todoID"])
	if err != nil {
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

// TodoCreate to create a new todo
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1*megaBytes))
	// Protect for big object
	if err != nil {
		panic(err)
	}
	// Ensure the body is closed
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	// Check the input is in valid format
	if err := json.Unmarshal(body, &todo); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(jsonMsg{
			Code: http.StatusUnprocessableEntity,
			Msg:  "Cannot unmarshal JSON object into struct",
		}); err != nil {
			panic(err)
		}
		return
	}

	// Create todo in database
	t := RepoCreateTodo(todo)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
