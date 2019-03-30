package generalization

import (
	"bitbucket.org/dargzero/k-anon/hierarchy"
	"bitbucket.org/dargzero/k-anon/partition"
)

// HierarchyGeneralizer is an implementation of Generalizer which uses a
// generalization Hierarchy to calculate generalized values.
type HierarchyGeneralizer struct {
	Hierarchy hierarchy.Hierarchy
}

// Generalizes the given partition to level (n-l) of the Hierarchy, where n is the
// total number of levels in the Hierarchy and l is the level of p in the Hierarchy.
// Consequently for leaf partitions this method will always generalize them to level n.
func (g *HierarchyGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	if n >= g.Levels() {
		return nil
	}
	h := g.Hierarchy.Find(p)
	if h == nil {
		return nil
	}
	l := h.Levels() - 1
	for i := l; i < n; i++ {
		h = h.Parent()
	}
	return h.Partition()
}

func (g *HierarchyGeneralizer) Levels() int {
	return g.Hierarchy.Levels()
}

func (g *HierarchyGeneralizer) InitItem(item interface{}) partition.Partition {
	return partition.NewSet(item)
}
