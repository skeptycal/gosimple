package structures

type (
	// BinarySearchTree is a tree structure where three restrictions
	// are in place to improve performance:
	//
	// 1. A Node can have two children at most.
	//
	// 2. For any given parent node, the child on the left has a
	// value less or equal to the parent. For any given parent node,
	// the child on the right has a value greater or equal to the parent.
	//
	// 3. No two nodes can have the same value.
	//
	// The biggest advantage of BST's is that we can search through
	// them in logarithmic time. Extremely popular for storing large
	// quantities of data that must be easily searchable.
	BinarySearchTree struct {
		root   *vertex // parent = nil
		height int     // number of edges on the longest path
	}

	BSTvertex struct {
		parent *vertex
		left   *vertex // values are less than current vertex
		right  *vertex // values are greater than current vertex
		depth  int     // number of edges between vertex and root node
	}
)
