package binarysearchtree

import (
	"fmt"
	"io"
	"strings"
)

// Min returns the item with min value stored in the tree
func (bst *Bst[K, V]) Min() *V {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil // TODO: should this be an error?
		// panic("cannot call Min() on empty tree")
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max returns the item with max value stored in the tree
func (bst *Bst[K, V]) Max() *V {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil // TODO: should this be an error?
		// panic("cannot call Max() on empty tree")
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// String returns a CLI readable rendering of the tree
func (bst *Bst[K, V]) String() string {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	sb := &strings.Builder{}
	fmt.Fprintln(sb, "------------------------------------------------")
	bst.stringify(sb, bst.root, 0)
	fmt.Fprintln(sb, "------------------------------------------------")
	return sb.String()
}

func (bst *Bst[K, V]) stringify(w io.Writer, n *Node[K, V], level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		bst.stringify(w, n.left, level)
		fmt.Fprintf(w, format+"%d\n", n.key)
		bst.stringify(w, n.right, level)
	}
}
