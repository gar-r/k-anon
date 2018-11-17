package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func BuildAnonGraph(costGraph graph.WeightedUndirected, k int) graph.Directed {
	g := simple.NewDirectedGraph()
	for !isComplete(g, k) {

	}
	return g
}

func isComplete(g graph.Directed, k int) bool {
	if isEmpty(g) {
		return false
	}
	components := topo.ConnectedComponents(graph.Undirect{G: g})
	for _, c := range components {
		if len(c) < k {
			return false
		}
	}
	return true
}

func isEmpty(g graph.Directed) bool {
	return g.Nodes() == nil || g.Nodes().Len() < 1
}
