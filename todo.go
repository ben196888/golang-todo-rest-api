package main

import (
	"encoding/json"
	"time"
)

// Todo item structure
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// MarshalBinary to marshal sturcture Todo to json format Todo
func (t *Todo) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

// UnmarshalBinary to unmarshal json format Todo to structure Todo
func (t *Todo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &t)
}

// Todos type is list of todo
type Todos []Todo
