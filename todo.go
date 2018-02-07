package main

import "time"

// Todo item structure
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos type is list of todo
type Todos []Todo
