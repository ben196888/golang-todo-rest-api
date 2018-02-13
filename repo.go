package main

import "fmt"

var currentID int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// RepoFindTodo find todo by id, returns empty todo if not found
func RepoFindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	// return empty todo if not found
	return Todo{}
}

// RepoCreateTodo append a todo w/ auto-increment id
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo remove a todo by id, return error if todo not existed
func RepoDestroyTodo(id int) error {
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo w/ id of %d to delete", id)
}
