package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

func decompose(g graph.Undirected, k int) {
	for {
		c := pickComponentToSplit(g, k)
		if c == nil {
			break
		}
	}
}

func pickComponentToSplit(g graph.Undirected, k int) []graph.Node {
	components := topo.ConnectedComponents(g)
	threshold := getThreshold(k)
	for _, c := range components {
		if len(c) > threshold {
			return c
		}
	}
	return nil
}

func getThreshold(k int) int {
	threshold := 2*k - 1
	if 3*k-5 > threshold {
		threshold = 3*k - 5
	}
	return threshold
}
