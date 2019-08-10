package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
)

// RangeGeneralizer is a Generalizer which works with ranges.
type RangeGeneralizer struct {
	r partition.Range
}

// NewIntRangeGeneralizer creates a new RangeGeneralizer for integers.
func NewIntRangeGeneralizer(min, max int) *RangeGeneralizer {
	return &RangeGeneralizer{
		r: partition.NewIntRange(min, max),
	}
}

// NewFloatRangeGeneralizer creates a new RangeGeneralizer for floats.
func NewFloatRangeGeneralizer(min, max float64) *RangeGeneralizer {
	return &RangeGeneralizer{
		r: partition.NewFloatRange(min, max),
	}
}

// Generalize generalizes the partition n levels further and returns the resulting partition.
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

// Levels returns the number of levels in the hierarchy.
func (g *RangeGeneralizer) Levels() int {
	return g.r.MaxSplit() + 1
}

// InitItem initializes an item into a partition.
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
