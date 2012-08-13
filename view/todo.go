package view

import (
	"github.com/yanatan16/go-todo-app/model"
	"html/template"
	"net/http"
)

// Todo is a view (http.Handler) that serves the todo app.
type Todo struct {
	*template.Template
}

// Create a Todo view.
// Arguments:
//	root - Template document root
func NewTodo(root string) Todo {
	t := Todo{
		template.Must(template.ParseFiles(root + "/todo.html")),
	}
	return Todo(t)
}

// ServeHTTP function lets Todo view implement http.Handler
func (t Todo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	t.Execute(res, model.List)
}
