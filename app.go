package main

import (
	"flag"
	"fmt"
	"github.com/yanatan16/go-todo-app/controller"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/view"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port             int
	prod             bool
	templateRoot     string
	serveStaticFiles bool
)

func Start() {
	mux := http.NewServeMux()

	// Instantiate an http.Server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Model
	model.Init(prod)

	// View
	view.Init(mux, templateRoot)

	// Controller
	controller.Init(mux)

	// Static files (non-prod)
	if serveStaticFiles {
		ServeStatic("./static", mux)
	}

	log.Printf("Now starting Todo App Server on port %d...", port)
	log.Fatal(server.ListenAndServe())
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

func ServeStatic(root string, mux *http.ServeMux) {
	mux.Handle("/static", http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			fn := root + "/" + strings.TrimLeft(req.URL.Path, "/static")
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				res.WriteHeader(404)
				res.Write([]byte("File not found!"))
				log.Println("Could not read file!", err)
				return
			}

			_, err = res.Write(data)
			if err != nil {
				log.Println("Could not write data back.", fn, err)
			}

		}))
}
