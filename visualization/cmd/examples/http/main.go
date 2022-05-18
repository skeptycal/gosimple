package main

import (
	"net/http"

	"github.com/skeptycal/gosimple/cli/errorlogger"
	"github.com/skeptycal/gosimple/visualization/cmd/examples/http/handlers"
)

const addr = "localhost:8088"

var (
	log = errorlogger.New()
)

func main() {
	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/headers", handlers.Headers)
	http.HandleFunc("/graph", handlers.Graph)

	// Where the magic happens
	// f, _ := os.Create("bar.html")
	// bar.Render(f)
	http.ListenAndServe(addr, nil)
}
