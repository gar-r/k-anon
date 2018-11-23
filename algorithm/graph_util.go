package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"math/rand"
	"time"
)

func pickRandomVertex(component []graph.Node) graph.Node {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(component))
	return component[i]
}

func containsNode(component []graph.Node, u graph.Node) bool {
	for _, n := range component {
		if n == u {
			return true
		}
	}
	return false
}
