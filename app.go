package main

import (
	"flag"
	"fmt"
	"github.com/yanatan16/go-todo-app/controller"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/view"
	"log"
	"net/http"
	"time"
)

var (
	port         int
	templateRoot string
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
	model.Init()

	// View
	view.Init(mux, templateRoot)

	// Controller
	controller.Init(mux)

	log.Printf("Now starting Todo App Server...")
	log.Fatal(server.ListenAndServe())
}

func main() {

	// Parse out the command line arguments
	flag.IntVar(&port, "port", 8080, "Port to listen on.")
	flag.StringVar(&templateRoot, "root", "./view/templates", "Template file root directory")
	flag.Parse()

	// Run the server
	Start()
}
