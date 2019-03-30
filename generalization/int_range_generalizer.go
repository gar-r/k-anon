package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"math"
)

type IntRangeGeneralizer struct {
	min, max int
	maxRange *partition.IntRange
}

func NewIntRangeGeneralizer(min, max int) *IntRangeGeneralizer {
	return &IntRangeGeneralizer{
		min:      min,
		max:      max,
		maxRange: partition.NewIntRange(min, max),
	}
}

func (g *IntRangeGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	_, success := p.(*partition.IntRange)
	if !success {
		return nil
	}
	path := make([]partition.IntRange, 0)
	g.trace(p, g.maxRange, &path)
	maxLevels := len(path) - 1
	level := indexOf(path, p)
	if level == -1 || n > maxLevels {
		return nil
	}
	if n <= level {
		return p
	}
	return &path[n]
}

func (g *IntRangeGeneralizer) Levels() int {
	return int(math.Ceil(
		math.Log2(float64(g.max-g.min+1))) + 1)
}

func (g *IntRangeGeneralizer) InitItem(item interface{}) partition.Partition {
	value, _ := item.(int)
	return partition.NewIntRange(value, value)
}

func (g *IntRangeGeneralizer) trace(p partition.Partition, r *partition.IntRange, path *[]partition.IntRange) {
	prepend(r, path)
	if !r.CanSplit() {
		for len(*path) < g.Levels() { // pad path to max level if needed
			prepend(r, path)
		}
		return
	}
	r1, r2 := r.Split()
	if r1.ContainsPartition(p) {
		g.trace(p, r1, path)
	} else {
		g.trace(p, r2, path)
	}
}

func prepend(item *partition.IntRange, path *[]partition.IntRange) {
	temp := make([]partition.IntRange, 1)
	temp[0] = *item
	*path = append(temp, *path...)
}

func indexOf(path []partition.IntRange, p partition.Partition) int {
	for i, q := range path {
		if p.Equals(&q) {
			return i
		}
	}
	return -1
}
