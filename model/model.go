// The model represents the core business logic
// for the todo list web app.
package model

import "github.com/yanatan16/go-todo-app/model/todo"

var List *todo.List

// Initialize the app's model.
func Init() {
	List = todo.NewList()
}
