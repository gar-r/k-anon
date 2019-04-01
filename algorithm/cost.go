package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"bitbucket.org/dargzero/k-anon/partition"
	"errors"
	"fmt"
)

func CalculateCost(r1, r2 *model.Row, schema *model.Schema) (float64, error) {
	var cost float64
	for j, col := range schema.Columns {
		if col.IsIdentifier() {
			d1 := r1.Data[j]
			d2 := r2.Data[j]
			fraction, err := calculateCostFraction(d1, d2, col.GetGeneralizer())
			if err != nil {
				return 0, err
			}
			cost += fraction * col.GetWeight()
		}
	}
	return cost, nil
}

func calculateCostFraction(p1, p2 partition.Partition, g generalization.Generalizer) (float64, error) {
	maxLevels := g.Levels()
	for level := 0; level < maxLevels; level++ {
		g1 := g.Generalize(p1, level)
		g2 := g.Generalize(p2, level)
		if g1 != nil && g1.Equals(g2) {
			return float64(level) / float64(maxLevels-1), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("data cannot be generalized into same partition: %v, %v", p1, p2))
}
