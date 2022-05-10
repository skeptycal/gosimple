package structures

type (
	vertex struct {
		parent   *vertex
		children []*vertex
		depth    int
	}

	edge struct {
		first          *vertex
		second         *vertex
		directionality int
	}

	Tree struct {
		root   *vertex // parent = nil
		height int     // number of edges on the longest path
	}
)
