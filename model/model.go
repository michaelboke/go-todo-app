// The model represents the core business logic
// for the todo list web app.
package model

import (
	"github.com/yanatan16/go-todo-app/model/todo"
)

// The global model object
var App *TodoApp

// TodoApp represents the entire model for the Todo App
type TodoApp struct {
	// Production Flag
	Prod bool
	// Actual todo list model
	List *todo.List
}

// Create a new Todo app model.
func NewTodoApp(prod bool) *TodoApp {
	return &TodoApp{
		Prod: prod,
		List: todo.NewList(),
	}
}

// Initialize the app's model.
func Init(prod bool) {
	App = NewTodoApp(prod)
}
