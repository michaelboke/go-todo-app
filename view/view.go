// The views represent the presentation layer of 
// the todo web app.
package view

import (
	"github.com/yanatan16/go-todo-app/model"
	"net/http"
)

// Initalize the app's views.
// Arguments:
//	svr - http server mux to listen on
//  template_root - Root directory for templates
func Init(svr *http.ServeMux, templateRoot string) {
	t := NewTodo(model.App, templateRoot)

	svr.Handle("/", http.Handler(t))
}
