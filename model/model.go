// The model represents the core business logic
// for the todo list web app.
package model

import (
	"github.com/hoisie/web"
	"github.com/yanatan16/go-todo-app/helpers"
	"github.com/yanatan16/go-todo-app/model/todo"
)

// The global model object
var App *TodoApp

// TodoApp represents the entire model for the Todo App
type TodoApp struct {
	// Actual todo list model
	List *todo.List
}

// Create a new Todo app model.
func NewTodoApp() *TodoApp {
	return &TodoApp{
		List: todo.NewList(),
	}
}

// Initialize the app's model.
func Init(svr *web.Server) {
	App = NewTodoApp()

	helpers.BindController(svr, "/todo/list", App.List)
}
