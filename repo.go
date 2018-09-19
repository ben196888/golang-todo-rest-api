package main

import (
	"fmt"
	"log"
)

var todos Todos

func init() {
	log.Println("Establish redis client...")
	redisClient = NewRedisClient()
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// RepoFindTodo find todo by id, returns empty todo if not found
func RepoFindTodo(id int) Todo {
	var t Todo
	val, err := redisClient.Get(string(id)).Result()
	if err != nil {
		fmt.Println(err)
		// return empty todo if not found
		return Todo{}
	}
	err = t.UnmarshalBinary([]byte(val))
	if err != nil {
		fmt.Println(err)
		return Todo{}
	}
	return t
}

// RepoCreateTodo append a todo w/ auto-increment id
func RepoCreateTodo(t Todo) Todo {
	todos = append(todos, t)
	curID, err := redisClient.Incr("todoCounter").Result()
	if err != nil {
		panic(err)
	}
	t.ID = int(curID)
	serr := redisClient.Set(string(curID), &t, 0).Err()
	if serr != nil {
		panic(serr)
	}
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
