// The controllers represent the actions that can be
// taken with the web application.
package controller

import (
	"encoding/json"
	"github.com/yanatan16/go-todo-app/model"
	"net/http"
	"net/url"
)

// Initialize the app's controllers
// Arguments:
// 	svr - Http ServeMux to handle on
func Init(svr *http.ServeMux) {
	svr.Handle("/add", Controller(add))
	svr.Handle("/done", Controller(done))
}

// An Controller on a list, triggered by a request
// Arguments:
//	List - A list to operate on
//	Values - Request arguments
// Returns:
//	map - Response data to be encoded as JSON
//	error - An error, if applicable
type Controller func(*model.TodoApp, url.Values) (map[string]interface{}, error)

// Make an http handler from a controller
func (ctrl Controller) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data, err := ctrl(model.App, req.URL.Query())
	if err != nil {
		// Respond with an error
		res.WriteHeader(400)
		res.Write([]byte(err.Error()))
	}

	// Write the response
	if data != nil {
		// Automatically write code 200
		enc := json.NewEncoder(res)
		if err = enc.Encode(data); err != nil {
			res.WriteHeader(500)
			res.Write([]byte(err.Error()))
		}
	} else {
		// No response included
		res.WriteHeader(204)
	}
}
