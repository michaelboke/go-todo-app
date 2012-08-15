// The views represent the presentation layer of 
// the todo web app.
package view

import (
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
	"github.com/yanatan16/go-todo-app/model"
	"log"
)

// Initalize the app's views.
// Arguments:
//	svr - http server mux to listen on
//  template_root - Root directory for templates
func Init(svr *web.Server, templateRoot string) {
	svr.Get("/",
		TemplateHandler(
			templateRoot+"/todo.html.mustache",
			model.App))
}

// Create a web handler using a mustache template
// Arguments:
//	fn is a file name
// 	data is the data that will be used to render the template
// Returns: Web handler
func TemplateHandler(fn string, data interface{}) func(*web.Context) {
	t, err := mustache.ParseFile(fn)
	if err != nil {
		log.Fatalf("Error parsing %s: %s", fn, err.Error())
	}

	return func(ctx *web.Context) {
		ctx.WriteString(t.Render(data))
	}
}
