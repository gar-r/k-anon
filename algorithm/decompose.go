package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"math"
	"math/rand"
	"time"
)

func decompose(g *simple.UndirectedGraph, k int) {
	for {
		c := pickComponentToSplit(g, k)
		if c == nil {
			break
		}
		partitionComponent(g, c, k)
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

func partitionComponent(g *simple.UndirectedGraph, component []graph.Node, k int) {
	u, v, t := getSplitParams(g, component, k)
	s := len(component)
	if t >= k && s-t >= k {
		splitTypeA(g, u, v)
	} else if s-t == k-1 {
		splitTypeB(g, u, v)
	} else if t == k-1 {
		splitTypeC(g, u, v)
	} else {
		var comp1, comp2 []graph.Node
		subTrees := getSubTrees(g, u)
		for _, subTree := range subTrees {
			if len(comp1) < k-1 {
				comp1 = append(comp1, subTree...)
			} else {
				comp2 = append(comp2, subTree...)
			}
		}
		if len(comp1) == k-1 {
			cutEdgesToComponent(g, comp2, u)
			sv := createSteinersVertex(g)

		} else if len(comp2) == k-1 {
			cutEdgesToComponent(g, comp1, u)
		} else {

		}
	}
}

func cutEdgesToComponent(g simple.UndirectedGraph, component []graph.Node, node graph.Node) {
	for _, n := range component {
		g.RemoveEdge(node.ID(), n.ID())
	}
}

func createSteinersVertex(g *simple.UndirectedGraph) graph.Node {
	sv := g.NewNode() // TODO: mark Steiner's vertex
	g.AddNode(sv)
	return sv
}

func splitTypeA(g *simple.UndirectedGraph, u graph.Node, v graph.Node) {
	g.RemoveEdge(u.ID(), v.ID())
}

func splitTypeB(g *simple.UndirectedGraph, u graph.Node, v graph.Node) {
	sv := createSteinersVertex(g)
	edges := g.From(v.ID())
	for edges.Next() {
		n := edges.Node()
		if u.ID() != n.ID() {
			g.RemoveEdge(v.ID(), n.ID())
		}
		g.NewEdge(sv, n)
	}
}

func splitTypeC(g *simple.UndirectedGraph, u graph.Node, v graph.Node) {
	splitTypeB(g, v, u)
}

func getSplitParams(g *simple.UndirectedGraph, component []graph.Node, k int) (graph.Node, graph.Node, int) {
	u := pickRandomVertex(component)
	for {
		largest := getLargestComponent(getSubTrees(g, u))
		v := getNextRoot(largest, g, u)
		if len(component)-len(largest) >= k-1 {
			return u, v, len(largest)
		}
		u = v
	}
}

func getSubTrees(g graph.Undirected, root graph.Node) [][]graph.Node {
	gCopy := simple.NewUndirectedGraph()
	graph.Copy(gCopy, g)
	gCopy.RemoveNode(root.ID())
	return topo.ConnectedComponents(gCopy)
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
