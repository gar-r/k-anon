package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
)

type RangeGeneralizer struct {
	r partition.Range
}

func NewIntRangeGeneralizer(min, max int) *RangeGeneralizer {
	return &RangeGeneralizer{
		r: partition.NewIntRange(min, max),
	}
}

func NewFloatRangeGeneralizer(min, max float64) *RangeGeneralizer {
	return &RangeGeneralizer{
		r: partition.NewFloatRange(min, max),
	}
}

func (g *RangeGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	_, success := p.(partition.Range)
	if !success {
		return nil
	}
	path := make([]partition.Range, 0)
	g.trace(p, g.r, &path)
	maxLevel := g.Levels() - 1
	level := indexOf(path, p)
	if level == -1 || n > maxLevel {
		return nil
	}
	if n <= level {
		return p
	}
	return path[n]
}

func (g *RangeGeneralizer) Levels() int {
	return g.r.MaxSplit() + 1
}

func (g *RangeGeneralizer) InitItem(item interface{}) partition.Partition {
	return g.r.InitItem(item)
}

func (g *RangeGeneralizer) trace(p partition.Partition, r partition.Range, path *[]partition.Range) {
	prepend(r, path)
	if !r.CanSplit() {
		levels := g.Levels()
		for len(*path) < levels { // pad path to max level if needed
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

func prepend(item partition.Range, path *[]partition.Range) {
	temp := make([]partition.Range, 1)
	temp[0] = item
	*path = append(temp, *path...)
}

func indexOf(path []partition.Range, p partition.Partition) int {
	for i, q := range path {
		if p.Equals(q) {
			return i
		}
	}
	return -1
}
