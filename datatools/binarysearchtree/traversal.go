package binarysearchtree

// InOrderTraverse visits all nodes with in-order traversing
func (bst *Bst[K, V]) InOrderTraverse(f func(V)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.inOrderTraverse(bst.root, f)
}

func (bst *Bst[K, V]) inOrderTraverse(n *Node[K, V], f func(V)) {
	if n != nil {
		bst.inOrderTraverse(n.left, f)
		f(n.value)
		bst.inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *Bst[K, V]) PreOrderTraverse(f func(V)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	bst.preOrderTraverse(bst.root, f)
}

func (bst *Bst[K, V]) preOrderTraverse(n *Node[K, V], f func(V)) {
	if n != nil {
		f(n.value)
		bst.preOrderTraverse(n.left, f)
		bst.preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *Bst[K, V]) PostOrderTraverse(f func(V)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	bst.postOrderTraverse(bst.root, f)
}

func (bst *Bst[K, V]) postOrderTraverse(n *Node[K, V], f func(V)) {
	if n != nil {
		bst.postOrderTraverse(n.left, f)
		bst.postOrderTraverse(n.right, f)
		f(n.value)
	}
}
