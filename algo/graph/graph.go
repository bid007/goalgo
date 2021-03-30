package graph

const (
	WHITE = 1
	GREY  = 2
	BLACK = 3
)

// Vertex is a graph vertex
type Vertex struct {
	color int
	val   interface{}
	edges []*Vertex
}

// New returns a new instance of a graph vertex
func New(val interface{}) *Vertex {
	return &Vertex{
		color: WHITE,
		val:   val,
		edges: make([]*Vertex, 1),
	}
}

// UndirectedEdge creates an undirected connection (v1 <--> v2) between two vertices
func (v1 *Vertex) UndirectedEdge(v2 *Vertex) {
	v1.edges = append(v1.edges, v2)
	v2.edges = append(v2.edges, v1)
}

// DirectedEdge creates a directed connection (v1 ->v2) between two vertices
func (v1 *Vertex) DirectedEdge(v2 *Vertex) {
	v1.edges = append(v1.edges, v2)
}

// MarkDiscovered marks the vertex as discovered during exploration
func (v1 *Vertex) MarkDiscovered() {
	v1.color = GREY
}

// IsDiscovered returns true if the vertex is already discovered during exploration
func (v1 *Vertex) IsDiscovered() bool {
	return v1.color == GREY
}

// MarkVisited marks the vertex as visited during exploration
func (v1 *Vertex) MarkVisited() {
	v1.color = BLACK
}

// IsVisited returns true if the vertex is already visited during exploration
func (v1 *Vertex) IsVisited() bool {
	return v1.color == BLACK
}

// HasCycle return true if there is a cycle in the graph
func (v *Vertex) HasCycle() bool {
	var (
		cycle bool
		dfs   func(v *Vertex)
	)

	dfs = func(v *Vertex) {
		if cycle {
			return
		}
		v.MarkDiscovered()
		for _, vertex := range v.edges {
			if vertex == nil || vertex.IsVisited() {
				continue
			}
			if vertex.IsDiscovered() {
				cycle = true
				return
			}
			dfs(vertex)
		}
		v.MarkVisited()
	}
	dfs(v)
	return cycle
}

// Reset changes state of all the vertices of the graph to unexplored and unvisited i.e. WHITE
func (v *Vertex) Reset() {
	queue := []*Vertex{v}
	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]
		if vertex == nil {
			continue
		}
		vertex.color = WHITE
		queue = append(queue, vertex.edges...)
	}
}
