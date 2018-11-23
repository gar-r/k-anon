package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
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

func getSubTrees(g graph.Undirected, root graph.Node) [][]graph.Node {
	gCopy := simple.NewUndirectedGraph()
	graph.Copy(gCopy, g)
	gCopy.RemoveNode(root.ID())
	return topo.ConnectedComponents(gCopy)
}
