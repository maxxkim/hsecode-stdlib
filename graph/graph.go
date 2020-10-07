package graph

type Type int
type Node interface {
	ID() int
}
type Graph struct {
	Type  Type
	nodes map[int]Node
	edges map[int]map[int]interface{}
}


const (
	Directed   Type = 0
	Undirected Type = 1
)

func New(graphType Type) *Graph {
	G := &Graph{
		Type:  graphType,
		nodes: make(map[int]Node),
		edges: make(map[int]map[int]interface{}),
	}
	return G
}

func (g *Graph) AddNode(node Node) {
	id := node.ID()
	g.nodes[id] = node
}

func (g *Graph) AddEdge(u, v int, edgeData interface{}) {
	if g.nodes[u] == nil || g.nodes[v] == nil {
		panic("node does not exist")
	}
	child, ok := g.edges[u]
	if !ok {
		child = make(map[int]interface{})
		g.edges[u] = child
	}
	child[v] = edgeData
}

func (g *Graph) Edge(u, v int) (interface{}, bool) {
	if g.Type == Directed {
		if val1, ok := g.edges[u]; ok {
			if val2, ok := val1[v]; ok {
				return val2, true
			}
		}
	} else if g.Type == Undirected {
		if val1, ok := g.edges[u]; ok {
			if val2, ok := val1[v]; ok {
				return val2, true
			}
		}
		if val1, ok := g.edges[v]; ok {
			if val2, ok := val1[u]; ok {
				return val2, true
			}
		}
	}
	return nil, false
}

func (g *Graph) Node(id int) (Node, bool) {
	if data := g.nodes[id]; data != nil {
		return data, true
	} else {
		return data, false
	}
}

func (g *Graph) Nodes(f func(Node)) {
	for _, value := range g.nodes {
		f(value)
	}
}

func (g *Graph) Edges(f func(u, v Node, edgeData interface{})) {
	for i := range g.edges {
		for k := range g.edges[i] {
			f(g.nodes[i], g.nodes[k], g.edges[i][k])
		}
	}
}

func (g *Graph) Neighbours(u int, f func(v Node, edgeData interface{})) {
	if g.nodes[u] == nil {
		panic("node does not exist")
	}
	for i := range g.nodes {
		if data, ok := g.Edge(u, i); ok {
			f(g.nodes[i], data)
		}
	}
}
