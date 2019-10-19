package db

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Todo struct {
	Id          int
	Title       string
	Description string
}

var storageFile string

var todos []Todo

func Init() {
	home, _ := os.UserHomeDir()

	storageFile = filepath.Join(home, ".gotodos")
}

func GetUniqueId(seed int) int {
	var matches []bool

	for _, todo := range todos {
		if todo.Id == seed {
			matches = append(matches, true)
		}
	}

	if len(matches) == 0 {
		return seed
	} else {
		return GetUniqueId(seed + 1)
	}
}

func ReadFromFile() {
	Init()

	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		os.Create(storageFile)
	}

	rawTodos, _ := ioutil.ReadFile(storageFile)

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

func WriteToFile() {
	Init()

	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		os.Create(storageFile)
	}

	var todosToWriteToFile string

	for _, todo := range todos {
		todosToWriteToFile += strconv.Itoa(todo.Id) + "," + todo.Title + "," + todo.Description + "\n"
	}

	ioutil.WriteFile(storageFile, []byte(todosToWriteToFile), 0400)
}

func Create(newTodo Todo) {
	todos = append(todos, newTodo)
}

func Read() []Todo {
	return todos
}

func Update(todoToBeUpdated Todo) (Todo, error) {
	var newTodos []Todo
	var updatedTodo Todo

	for _, todo := range todos {
		if todo.Id == todoToBeUpdated.Id {
			var updatedTodoTitle string
			var updatedTodoDescription string

			if todoToBeUpdated.Title == "" {
				updatedTodoTitle = todo.Title
			} else {
				updatedTodoTitle = todoToBeUpdated.Title
			}

			if todoToBeUpdated.Description == "" {
				updatedTodoDescription = todo.Description
			} else {
				updatedTodoDescription = todoToBeUpdated.Description
			}

			updatedTodo = Todo{Id: todoToBeUpdated.Id, Title: updatedTodoTitle, Description: updatedTodoDescription}

			newTodos = append(newTodos, updatedTodo)

		} else {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos

	if updatedTodo.Title == "" {
		return updatedTodo, errors.New("Todo with ID " + strconv.Itoa(todoToBeUpdated.Id) + " could not be found")
	} else {
		return updatedTodo, nil
	}
}

func Delete(id int) {
	var newTodos []Todo

	for _, todo := range todos {
		if !(todo.Id == id) {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos
}
