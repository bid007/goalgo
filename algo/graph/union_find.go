package graph

// UnionFind represents structure that support union and find on sets
type UnionFind struct {
	parent map[int]int
	size   map[int]int
}

// NewSet creates a new instance of union find with given elems
func NewSet(sets ...int) *UnionFind {
	uf := &UnionFind{
		parent: make(map[int]int),
		size:   make(map[int]int),
	}
	for _, s := range sets {
		uf.parent[s] = s
		uf.size[s] = 1
	}
	return uf
}

// Find returns the parent of a set
func (uf *UnionFind) Find(s int) int {
	if uf.parent[s] == s {
		return s
	}
	return uf.Find(uf.parent[s])
}

// Union merges two set
func (uf *UnionFind) Union(s1, s2 int) {
	var (
		p1 = uf.Find(s1)
		p2 = uf.Find(s2)
	)
	if p1 != p2 {
		if uf.size[p1] <= uf.size[p2] {
			uf.parent[p1] = p2
			uf.size[p2] += uf.size[p1]
		} else {
			uf.parent[p2] = p1
			uf.size[p1] += uf.size[p2]
		}
	}
}
