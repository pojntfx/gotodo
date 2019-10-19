package todo

type Todo struct {
	id          int
	title       string
	description string
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
		if todo.id == todoToBeUpdated.id {
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
		if !(todo.id == todoToBeDeleted.id) {
			newTodos = append(newTodos, todo)
		}
	}

	todos = newTodos
}
