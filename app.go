package main

import (
	"flag"
	"fmt"
	"github.com/yanatan16/go-todo-app/controller"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/view"
	"io/ioutil"
	"log"
	"github.com/hoisie/web"
)

var (
	port             int
	prod             bool
	templateRoot     string
	serveStaticFiles bool
)

func Start() {
	server := web.NewServer()

	// Model
	model.Init(prod)

	// View
	view.Init(server, templateRoot)

	// Controller
	controller.Init(server)

	// Static files (non-prod)
	if serveStaticFiles {
		ServeStatic("./static", server)
	}

	log.Printf("Now starting Todo App Server on port %d...", port)
	server.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

func main() {

	// Parse out the command line arguments
	flag.IntVar(&port, "port", 8080, "Port to listen on.")
	flag.StringVar(&templateRoot, "root", "./view/templates", "Template root directory.")
	flag.BoolVar(&prod, "prod", false, "Production")
	flag.BoolVar(&serveStaticFiles, "static", false, "Serve Static files from Go")
	flag.Parse()

	// Correct conflicts
	serveStaticFiles = prod && serveStaticFiles // No serving static in prod
	if port == 8080 && prod {
		port = 80 // Default port is 80 in prod
	}

	// Run the server
	Start()
}

func ServeStatic(root string, svr *web.Server) {
	svr.Get("/static(.*)", 
		func(ctx *web.Context, path string) {
			fn := root + path
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				ctx.NotFound("File not found!")
				log.Println("Could not read file!", err)
				return
			}

			ctx.WriteString(string(data))
		})
}
