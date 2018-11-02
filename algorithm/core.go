package algorithm

import (
	"gonum.org/v1/gonum/graph/simple"
	"k-anon/model"
	"math"
)

func BuildCoreGraph(t *model.Table) *simple.WeightedUndirectedGraph {
	g := simple.NewWeightedUndirectedGraph(0, math.MaxFloat64)
	for i := range t.Rows {
		node := simple.Node(i)
		g.AddNode(node)
	}
	return g
}
