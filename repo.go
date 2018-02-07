package main

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
