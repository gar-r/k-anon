package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"k-anon/model"
	"math"
)

func BuildCostGraph(t *model.Table) graph.WeightedUndirected {
	g := buildEmptyCostGraph(t)
	addCosts(g, t)
	return g
}

func buildEmptyCostGraph(t *model.Table) *simple.WeightedUndirectedGraph {
	g := simple.NewWeightedUndirectedGraph(0, math.MaxFloat64)
	for i := range t.Rows {
		node := simple.Node(i)
		g.AddNode(node)
	}
	return g
}

func addCosts(g *simple.WeightedUndirectedGraph, t *model.Table) {
	nodes := len(t.Rows)
	for i := 0; i < nodes; i++ {
		for j := 0; j < nodes; j++ {
			if i != j {
				v1 := t.Rows[i]
				v2 := t.Rows[j]
				cost := CalculateCost(v1, v2)
				edge := g.NewWeightedEdge(g.Node(int64(i)), g.Node(int64(j)), cost)
				g.SetWeightedEdge(edge)
			}
		}
	}
}
