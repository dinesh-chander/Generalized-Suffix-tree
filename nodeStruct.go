package generalizedSuffixTree

type node struct {
	edgeLabel   string
	child_nodes map[byte]*node
	indexData   map[int][]int
	parent      *node
}
