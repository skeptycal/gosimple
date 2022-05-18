// Package binarysearchtree creates a BinarySearchTree data structure
// for the Key and Value types specified in in the constraints of the
// generic structs and interfaces.
//
// Specifically, the key type must be ordered and the value type may
// be any valid type.
//
// References:
// https://hackthedeveloper.com/golang-binary-search-tree/
// https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
package binarysearchtree

import (
	"fmt"
	"sync"

	"github.com/skeptycal/gosimple/types/constraints"
)

// New returns a new empty BinarySearchTree. Most features require one or more nodes to be inserted. The first node inserted will be the root node and all subsequent nodes will be positioned automatically, often changing the root node.
func New[K constraints.Ordered, V any]() BSTer[K, V] {
	return new(Bst[K, V])
}

type (
	BSTer[K constraints.Ordered, V any] interface {
		Insert(key K, value V)       // inserts the Item t in the tree
		Search(key K) bool           // returns true if the Item t exists in the tree
		Remove(key K) *Node[K, V]    // removes the Item t from the tree
		InOrderTraverse(f func(V))   //visits all nodes with in-order traversing
		PreOrderTraverse(f func(V))  //visits all nodes with pre-order traversing
		PostOrderTraverse(f func(V)) //visits all nodes with post-order traversing
		Min() *V                     //returns the Item with min value stored in the tree
		Max() *V                     //returns the Item with max value stored in the tree
		String() string              //prints a CLI readable rendering of the tree
		Root() *Node[K, V]           // returns the root node of the tree
	}

	// Node a single Node that composes the tree
	Node[K constraints.Ordered, V any] struct {
		key   K
		value V
		// parent *Node[K, V]
		left  *Node[K, V]
		right *Node[K, V]
	}

	// Bst is a generic binary search tree
	Bst[K constraints.Ordered, V any] struct {
		lock sync.RWMutex
		root *Node[K, V]

		// Reference: https://flaviocopes.com/golang-data-structure-binary-search-tree/
	}
)

// InOrder Traversal
func InOrder[K constraints.Ordered, V any](root *Node[K, V]) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Print(root.value, " ")
	InOrder(root.right)
}
