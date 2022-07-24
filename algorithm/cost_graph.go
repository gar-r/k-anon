package algorithm

import (
	"math"

	"git.okki.hu/garric/k-anon/model"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

// BuildCostGraph creates a weighted cost-graph from the table.
func BuildCostGraph(t *model.Table) (graph.WeightedUndirected, error) {
	g := buildEmptyCostGraph(t)
	err := addCosts(g, t)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func buildEmptyCostGraph(t *model.Table) *simple.WeightedUndirectedGraph {
	g := simple.NewWeightedUndirectedGraph(0, math.MaxFloat64)
	for i := range t.GetRows() {
		node := simple.Node(i)
		g.AddNode(node)
	}
	return g
}

func addCosts(g *simple.WeightedUndirectedGraph, t *model.Table) error {
	nodes := len(t.GetRows())
	for i := 0; i < nodes; i++ {
		for j := 0; j < nodes; j++ {
			if i != j {
				v1 := t.GetRows()[i]
				v2 := t.GetRows()[j]
				cost, err := CalculateCost(v1, v2, t.GetSchema())
				if err != nil {
					return err
				}
				edge := g.NewWeightedEdge(g.Node(int64(i)), g.Node(int64(j)), cost)
				g.SetWeightedEdge(edge)
			}
		}
	}
	return nil
}
