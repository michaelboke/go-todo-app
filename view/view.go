// The views represent the presentation layer of 
// the todo web app.
package view

import (
	"github.com/hoisie/web"
	"github.com/yanatan16/go-todo-app/helpers"
	"github.com/yanatan16/go-todo-app/model"
    "log"
    "encoding/json"
)

// Initalize the app's views.
// Arguments:
//	svr - http server mux to listen on
//  template_root - Root directory for templates
func Init(svr *web.Server, templateRoot string) {
	helpers.SetTemplateRoot(templateRoot)

	svr.Get("/todo", helpers.TemplateLayoutView(
		"/todo.html.mustache",
		"/base.html.mustache",
		func() interface{}{
            list := model.App.List.GetAll()
            bootstrap, err := json.Marshal(list)
            if err != nil {
                log.Printf("ERROR: Parsing GetAll():",err,list)
                return nil
            }

            return map[string]string{
                "title": "Todo App",
                "list": string(bootstrap),
            }
        }))
}
