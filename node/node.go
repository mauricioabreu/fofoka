package node

type Node struct {
	ID        int
	Neighbors []*Node
}

func New(id int) *Node {
	return &Node{
		ID:        id,
		Neighbors: []*Node{},
	}
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.Neighbors = append(n.Neighbors, neighbor)
}
