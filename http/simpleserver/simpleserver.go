package simpleserver

import (
	"fmt"
	"net/http"
)

func Server() *Server {
	// s := http.
	return nil
}

func root() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})
}
