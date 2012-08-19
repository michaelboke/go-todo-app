// The controllers represent the actions that can be
// taken with the web application.
package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"io/ioutil"
)

// A Controller for a corresponding Backbone Model object.
type Controller interface {
	// Create model object based on parameters
	// attr is a json-formatted string of attributes
	// Return a json-formattable object of all model attributes
	Create(attr string) (interface{}, error)
	// Read a model object back
	// ID may be empty string
	// Return a json-formattable object of all model attributes
	Read(id string) (interface{}, error)
	// Update a model object based on parameters. 
	// ID is required and will be non-empty
	// attr is a json-formatted string of attributes
	// Return a json-formattable object of updated model attributes
	// If no attributes other than the updated ones changed, it is acceptable to return nil
	Update(id, attr string) (interface{}, error)
	// Delete a model object.
	// ID is required and will be non-empty
	Delete(id string) error
}

type Context struct {
	*web.Context
}

func (ctx *Context) readBody() (string, error) {
	str, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func (ctx *Context) writeJson(m interface{}) {
	str, err := json.Marshal(m)
	if err != nil {
		ctx.Abort(500, "Error marshalling map: "+err.Error())
	}
	ctx.WriteString(string(str))
}

func (ctx *Context) writeError(err error) {
	ctx.WriteHeader(400)
	rv := map[string]string{
		"error": err.Error(),
	}
	str, err2 := json.Marshal(rv)
	if err2 != nil {
		ctx.Abort(500, fmt.Sprintf("Error mashalling error(%s): %s", err.Error(), err2.Error()))
	}
	ctx.WriteString(string(str))
}

func BindController(svr *web.Server, path string, ctrl Controller) {
	// Create
	svr.Post(path, func(wctx *web.Context) {
		ctx := &Context{wctx}
		body, err := ctx.readBody()
		if err != nil {
			ctx.writeError(err)
			return
		}

		ret, err := ctrl.Create(body)
		if err != nil {
			ctx.writeError(err)
			return
		}

		ctx.writeJson(ret)
	})

	// Read
	svr.Get(path+"/?(.*)", func(wctx *web.Context, id string) {
		ctx := &Context{wctx}
		ret, err := ctrl.Read(id)
		if err != nil {
			ctx.writeError(err)
			return
		}

		ctx.writeJson(ret)
	})

	// Update
	svr.Put(path+"/(.+)", func(wctx *web.Context, id string) {
		ctx := &Context{wctx}
		body, err := ctx.readBody()

		if err != nil {
			ctx.writeError(err)
			return
		}

		ret, err := ctrl.Update(id, body)
		if err != nil {
			ctx.writeError(err)
			return
		}

		// Accept nil responses
		if ret != nil {
			ctx.writeJson(ret)
		} else {
			ctx.NotModified()
		}
	})

	// Delete
	svr.Delete(path+"/(.+)", func(wctx *web.Context, id string) {
		ctx := &Context{wctx}
		err := ctrl.Delete(id)
		if err != nil {
			ctx.writeError(err)
			return
		}

		ctx.NotModified()
	})
}
