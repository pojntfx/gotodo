package db

type Todo struct {
	Id          int
	Title       string
	Description string
}

var todos []Todo

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
