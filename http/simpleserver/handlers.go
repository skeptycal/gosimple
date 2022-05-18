package simpleserver

import (
	"net/http"

	"github.com/skeptycal/gosimple/http/binarysearchtree"
)

func New(name string) HandlerLister {
	return &handlerList{}
}

type (
	HandlerLister interface {
		binarysearchtree.BSTer[int, *handler]
		Name() string
		Register() error
	}

	Handler interface {
		Register()
		// Name() string
		// Fn() http.HandlerFunc
		// Description() string
	}
	handler struct { // BST nodes
		name string
		fn   http.HandlerFunc
		// desc string // optional description
	}

	handlerList struct { // BST root node
		binarysearchtree.Bst[int, *handler]
		name string
	}
)

var handlerList1 = binarysearchtree.New[int, *handler]()

func (list *handlerList) Name() string { return list.name }

// func (list *handlerList) Len() int { return 0 } // TODO: Implement this
func (list *handlerList) Register() error {
	list.InOrderTraverse(func(h *handler) { h.Register() })
	return nil
}

// func (h *handler) Name() string         { return h.name }
// func (h *handler) Description() string  { return h.desc }
// func (h *handler) Fn() http.HandlerFunc { return h.fn }
func (h *handler) Register() { http.HandleFunc(h.name, h.fn) }
