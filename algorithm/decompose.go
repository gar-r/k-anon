package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"math"
	"math/rand"
	"time"
)

func decompose(g graph.Undirected, k int) {
	for {
		c := pickComponentToSplit(g, k)
		if c == nil {
			break
		}
		pickRootVertex(g, c, k)
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

func pickRootVertex(g graph.Undirected, component []graph.Node, k int) graph.Node {
	u := pickRandomVertex(component)
	for {
		largest := getLargestSubTree(g, u)
		if len(component)-len(largest) >= k-1 {
			return u
		}
		u = getNextRoot(largest, g, u)
	}
}

func getLargestSubTree(g graph.Undirected, root graph.Node) []graph.Node {
	gCopy := simple.NewUndirectedGraph()
	graph.Copy(gCopy, g)
	gCopy.RemoveNode(root.ID())
	subtrees := topo.ConnectedComponents(gCopy)
	return getLargestComponent(subtrees)
}

func getLargestComponent(components [][]graph.Node) []graph.Node {
	max := math.MinInt64
	var result []graph.Node
	for _, c := range components {
		if len(c) > max {
			max = len(c)
			result = c
		}
	}
	return result
}

func getNextRoot(largest []graph.Node, g graph.Undirected, u graph.Node) graph.Node {
	for _, v := range largest {
		if g.HasEdgeBetween(u.ID(), v.ID()) {
			return g.Node(v.ID())
		}
	}
	panic("no edge between root candidate and largest sub-tree")
}

func pickRandomVertex(component []graph.Node) graph.Node {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(component))
	return component[i]
}

func getThreshold(k int) int {
	threshold := 2*k - 1
	if 3*k-5 > threshold {
		threshold = 3*k - 5
	}
	return threshold
}
