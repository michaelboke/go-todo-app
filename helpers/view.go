package helpers

import (
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
	"log"
)

var root string

func SetTemplateRoot(dir string) {
	root = dir
}

// Create a web handler using a mustache template
// Arguments:
//  fn is a file name
//  data is the data that will be used to render the template
// Returns: Web handler
func TemplateView(fn string, data interface{}) func(*web.Context) {
	t, err := mustache.ParseFile(root + fn)
	if err != nil {
		log.Fatalf("Error parsing %s: %s", fn, err.Error())
	}

	return func(ctx *web.Context) {
		ctx.WriteString(t.Render(data))
	}
}

// Create a web handler using a mustache template with layout
// Arguments:
//  fn is a file name
//  data is the data that will be used to render the template
// Returns: Web handler
func TemplateLayoutView(fn, layoutFn string, data interface{}) func(*web.Context) {
	base, err := mustache.ParseFile(root + layoutFn)
	if err != nil {
		log.Fatalf("Error parsing %s: %s", layoutFn, err.Error())
	}

	t, err := mustache.ParseFile(root + fn)
	if err != nil {
		log.Fatalf("Error parsing %s: %s", fn, err.Error())
	}

	return func(ctx *web.Context) {
		ctx.WriteString(t.RenderInLayout(base, data))
	}
}
