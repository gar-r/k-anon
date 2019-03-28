package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
)

func CalculateCost(r1, r2 *model.Row, schema *model.Schema) float64 {
	var cost float64
	for j, col := range schema.Columns {
		if col.IsIdentifier() {
			d1 := r1.Data[j]
			d2 := r2.Data[j]
			cost += calculateCostFraction(d1, d2, col.Generalizer)
		}
	}
	return cost
}

func calculateCostFraction(p1, p2 *generalization.Partition, g generalization.Generalizer) float64 {
	maxLevels := g.Levels()
	for level := 0; level < maxLevels; level++ {
		g1 := g.Generalize(p1, level)
		g2 := g.Generalize(p2, level)
		if g1.Equals(g2) {
			return float64(level) / float64(maxLevels-1)
		}
	}
	panic("data cannot be generalized into same partition")
}
