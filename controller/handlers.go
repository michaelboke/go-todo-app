// The controllers represent the actions that can be
// taken with the web application.
package controller

import (
	"encoding/json"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/hoisie/web"
)

// Initialize the app's controllers
// Arguments:
// 	svr - Http ServeMux to handle on
func Init(svr *web.Server) {
	svr.Get("/add", MakeHandler(add))
	svr.Get("/done", MakeHandler(done))
}

// An Controller on a list, triggered by a request
// Arguments:
//	List - A list to operate on
//	Values - Request arguments
// Returns:
//	map - Response data to be encoded as JSON
//	error - An error, if applicable
type Controller func(*model.TodoApp, map[string]string) (interface{}, error)

// Make an web handler from a controller
func MakeHandler(ctrl Controller) func(*web.Context, string) {
	return func(ctx *web.Context, val string) {
		data, err := ctrl(model.App, ctx.Params)
		if err != nil {
			// Respond with an error
			ctx.NotFound(err.Error())
			return
		}

		// Write the response
		if data != nil {
			// Automatically write code 200
			enc, err := json.Marshal(data)
			if err != nil {
				ctx.Abort(500, err.Error())
				return
			}

			ctx.WriteString(string(enc))
		} else {
			// No response included
			ctx.NotModified()
		}
	}
}
