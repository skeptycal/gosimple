package binarysearchtree

// Insert inserts the item t in the tree
func (bst *Bst[K, V]) Insert(key K, value V) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node[K, V]{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		bst.insertNode(bst.root, n)
	}
}

func (bst *Bst[K, V]) insertNode(root, n *Node[K, V]) {
	if n.key < root.key {
		if root.left == nil {
			root.left = n
		} else {
			bst.insertNode(root.left, n)
		}
	} else {
		if root.right == nil {
			root.right = n
		} else {
			bst.insertNode(root.right, n)
		}
	}
}

// Search returns true if the item t exists in the tree
func (bst *Bst[K, V]) Search(key K) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return bst.search(bst.root, key)
}

func (bst *Bst[K, V]) search(n *Node[K, V], key K) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return bst.search(n.left, key)
	}
	if key > n.key {
		return bst.search(n.right, key)
	}
	return true
}

// Remove removes the item t from the tree
func (bst *Bst[K, V]) Remove(key K) *Node[K, V] {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return bst.remove(bst.root, key)
}

func (bst *Bst[K, V]) remove(node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = bst.remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = bst.remove(node.right, key)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = bst.remove(node.right, node.key)
	return node
}

func (bst *Bst[K, V]) Root() *Node[K, V] { return bst.root }
