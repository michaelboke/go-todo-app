// The model represents the core business logic
// for the todo list web app.
package model

import "github.com/yanatan16/go-todo-app/model/todo"

var App *TodoApp

// TodoApp represents the entire model for the Todo App
type TodoApp struct {
	List *todo.List
}

func NewTodoApp() *TodoApp {
	return &TodoApp{todo.NewList()}
}

// Initialize the app's model.
func Init() {
	App = NewTodoApp()
}
