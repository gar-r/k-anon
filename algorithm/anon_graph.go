package algorithm

import (
	"bitbucket.org/dargzero/k-anon/model"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"math"
)

func BuildAnonGraph(table *model.Table, k int) (graph.Directed, error) {
	costGraph, err := BuildCostGraph(table)
	if err != nil {
		return nil, err
	}
	g := buildEmptyAnonGraph(table)
	for {
		components := UndirectedConnectedComponents(g)
		c := pickComponentToExtend(components, k)
		if c == nil {
			break
		}
		u := pickSourceVertex(g, c)
		v := pickTargetVertex(g, c, u, costGraph)
		g.SetEdge(g.NewEdge(u, v))
	}
	return g, nil
}

func pickSourceVertex(g graph.Directed, component []graph.Node) graph.Node {
	for _, n := range component {
		outgoing := g.From(n.ID())
		if outgoing.Len() == 0 {
			return n
		}
	}
	panic("no vertex without outgoing edges in component")
}

func pickTargetVertex(g graph.Directed, component []graph.Node, u graph.Node, costGraph graph.WeightedUndirected) graph.Node {
	var targetVertex graph.Node
	minWeight := math.MaxFloat64
	nodes := costGraph.From(u.ID())
	for nodes.Next() {
		n := nodes.Node()
		w, _ := costGraph.Weight(u.ID(), n.ID())
		if !containsNode(component, n) && w < minWeight {
			minWeight = w
			targetVertex = g.Node(n.ID())
		}
	}
	return targetVertex
}

func pickComponentToExtend(components [][]graph.Node, k int) []graph.Node {
	for _, c := range components {
		if len(c) < k {
			return c
		}
	}
	return nil
}

func buildEmptyAnonGraph(t *model.Table) *simple.DirectedGraph {
	g := simple.NewDirectedGraph()
	for i := range t.GetRows() {
		node := simple.Node(i)
		g.AddNode(node)
	}
	return g
}
