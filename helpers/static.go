package helpers

import (
	"github.com/hoisie/web"
	"io/ioutil"
	"log"
)

// Web.go handler to serve static files
// This should not be used in production
func ServeStatic(root string, svr *web.Server) {
	svr.Get("/static(.*)",
		func(ctx *web.Context, path string) {
			fn := root + path
			log.Printf("Serving static file %s", fn)
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				ctx.NotFound("File not found!")
				log.Println("Could not read file!", err)
				return
			}

			ctx.WriteString(string(data))
		})
}
