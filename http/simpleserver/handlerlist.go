package simpleserver

import (
	"github.com/skeptycal/gosimple/datatools/binarysearchtree"
)

func New(name string) HandlerLister {
	return &handlerList{name: name}
}

type (
	HandlerLister interface {
		binarysearchtree.BSTer[int, *Handler]
		Name() string
		Register() error
	}

	handlerList struct { // BST root node
		binarysearchtree.Bst[int, *Handler]
		name string
	}
)

func (list *handlerList) Name() string { return list.name }

// func (list *handlerList) Len() int { return 0 } // TODO: Implement this

func (list *handlerList) Register() error {
	list.InOrderTraverse(func(h *Handler) { h.Register() })
	return nil
}
