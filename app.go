package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"github.com/hoisie/web"
	"github.com/yanatan16/go-todo-app/helpers"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/view"
	"log"
)

var (
	port             int
	templateRoot     string
	serveStaticFiles bool
	useCGI           bool
)

func Start() {
	server := web.NewServer()

	// Model
	model.Init(server)

	// View
	view.Init(server, templateRoot)

	// Static files (non-prod)
	if serveStaticFiles {
		helpers.ServeStatic("./static", server)
	}

	log.Printf("Now starting Todo App Server on port %d...", port)
	if useCGI {
		server.RunFcgi(fmt.Sprintf("0.0.0.0:%d", port))
	} else {
		server.Run(fmt.Sprintf("0.0.0.0:%d", port))
	}
}

func main() {

	// Parse out the command line arguments
	flag.IntVar(&port, "port", 8000, "Port to listen on.")
	flag.StringVar(&templateRoot, "root", "./view/templates", "Template root directory.")
	flag.BoolVar(&serveStaticFiles, "static", false, "Serve Static files from Go")
	flag.BoolVar(&useCGI, "cgi", false, "User FastCGI")
	flag.Parse()

	// Run the server
	Start()
}
