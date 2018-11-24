package algorithm

import "k-anon/model"

func CalculateCost(v1, v2 *model.Vector) float64 {
	var cost float64
	for i := range v1.Items {
		d1 := v1.Items[i]
		d2 := v2.Items[i]
		if d1.IsIdentifier() && d2.IsIdentifier() {
			cost += calculateCostFraction(d1, d2)
		}
	}
	return cost
}

func calculateCostFraction(d1 *model.Data, d2 *model.Data) float64 {
	maxLevels := d1.Levels()
	for level := 0; level < maxLevels; level++ {
		if d1.Generalize(level).Equals(d2.Generalize(level)) {
			return float64(level) / float64(maxLevels-1)
		}
	}
	panic("data cannot be generalized into same partition")
}
