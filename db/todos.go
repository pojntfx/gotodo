package db

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Id          int
	Title       string
	Description string
}

var todos []Todo

func ReadFromFile() {
	if _, err := os.Stat("/tmp/todos.csv"); os.IsNotExist(err) {
		os.Create("/tmp/todos.csv")
	}

	rawTodos, _ := ioutil.ReadFile("/tmp/todos.csv")

	todoLines := strings.Split(string(rawTodos), "\n")

	var todosFromFile []Todo

	for _, todoFromFile := range todoLines {
		todoParts := strings.Split(todoFromFile, ",")
		if len(todoParts) == 3 {
			ID, _ := strconv.ParseInt(todoParts[0], 0, 64)
			todo := Todo{Id: int(ID), Title: todoParts[1], Description: todoParts[2]}
			todosFromFile = append(todosFromFile, todo)
		}
	}

	todos = todosFromFile
}

func Create(newTodo Todo) {
	todos = append(todos, newTodo)
}

func Read() []Todo {
	return todos
}

func Update(todoToBeUpdated Todo) {
	var newTodos []Todo

	for _, todo := range todos {
		if todo.Id == todoToBeUpdated.Id {
			newTodos = append(newTodos, todoToBeUpdated)
		} else {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos
}

func Delete(todoToBeDeleted Todo) {
	var newTodos []Todo

	for _, todo := range todos {
		if !(todo.Id == todoToBeDeleted.Id) {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos
}
