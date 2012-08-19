// The views represent the presentation layer of 
// the todo web app.
package view

import (
	"github.com/hoisie/web"
	"github.com/yanatan16/go-todo-app/helpers"
	"github.com/yanatan16/go-todo-app/model"
)

// Initalize the app's views.
// Arguments:
//	svr - http server mux to listen on
//  template_root - Root directory for templates
func Init(svr *web.Server, templateRoot string) {
	helpers.SetTemplateRoot(templateRoot)

	svr.Get("/todo",
		helpers.TemplateLayoutView(
			"/todo.html.mustache",
			"/base.html.mustache",
			model.App))
}
