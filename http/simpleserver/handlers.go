package simpleserver

import (
	"net/http"
)

type (

	// HF is a type alias for http.HandlerFunc.
	//
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	HF = http.HandlerFunc // func(http.ResponseWriter, *http.Request)
)

func NewHandler(pattern string, h HF) *Handler {
	return &Handler{name: pattern, fn: h}
}

type (

	// Handler interface {
	// 	Register()
	// 	Name() string
	// 	Fn() HF
	// 	// Description() string
	// }

	// BST nodes
	Handler struct {
		name string
		fn   HF
		// desc string // optional description
	}
)

func (h *Handler) Name() string { return h.name }
func (h *Handler) Fn() HF       { return h.fn }
func (h *Handler) Register()    { http.HandleFunc(h.name, h.fn) }

// func (h *handler) Description() string  { return h.desc }
